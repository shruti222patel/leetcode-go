package scrapper

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
)

type LeetCodeScrapper struct {
	URL   string
	ctx   context.Context
	Debug bool
}

func NewLeetCodeScrapper(url string, debug bool) *LeetCodeScrapper {
	lcs := LeetCodeScrapper{
		URL:   url,
		Debug: debug,
	}

	return &lcs
}

// ScapeData scrapes relevant data off of leetcode
// To update xpaths, open chrome inspector, select element, copy `full xpath`
func (lcs *LeetCodeScrapper) ScrapeData() (*LeetCodeProblem, error) {
	lcp := LeetCodeProblem{}

	opts := chromedp.DefaultExecAllocatorOptions[:]
	if lcs.Debug {
		opts = append(opts,
			chromedp.Flag("headless", false),
			// chromedp.Flag("disable-gpu", false),
			// chromedp.Flag("enable-automation", false),
			// chromedp.Flag("disable-extensions", false),
		)
	}

	allocCtx, cancelAlloc := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancelAlloc()

	// create context
	ctx, cancelCtx := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancelCtx()

	// // Add timeout
	// if !lcs.Debug {
	// 	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	// 	defer cancel()
	// }

	lcs.ctx = ctx

	if err := chromedp.Run(ctx,
		chromedp.Emulate(device.IPadlandscape),
		chromedp.Navigate(lcs.URL),
		chromedp.WaitVisible(`/html/body/div[1]/div`),
	); err != nil {
		return nil, fmt.Errorf("Couldn't run chrome browser")
	}

	// Find name & number
	nameNodes, err := lcs.getNode("/html/body/div[1]/div/div[2]/div/div/div[1]/div/div[1]/div[1]/div/div[2]/div/div[1]/div[1]/text()")
	if err != nil {
		return nil, fmt.Errorf("Couldn't find name node")
	}
	lcp.Number = lcs.cleanScrapedString(nameNodes[0].Dump("", "", false))
	lcp.Name = lcs.cleanScrapedString(nameNodes[2].Dump("", "", false))
	fmt.Println("Finished scrapping Name & Number")

	// Find description
	// descNodes, _ := lcs.getNode("/html/body/div[1]/div/div[2]/div/div/div[1]/div/div[1]/div[1]/div/div[2]/div/div[2]/div/p//text()")
	descNodes, _ := lcs.getText("/html/body/div[1]/div/div[2]/div/div/div[1]/div/div[1]/div[1]/div/div[2]/div/div[2]/div/p//text()")
	if err != nil {
		return nil, fmt.Errorf("Couldn't find descp node")
	}
	// lcp.Description = lcs.cleanScrapedString(lcs.concatNodeStr(descNodes))
	// lcp.Description = lcs.cleanScrapedString(lcs.concatNodeStr(descNodes))
	fmt.Println(descNodes)
	fmt.Println("Finished scrapping Description")

	// Find explaination
	exampleNodes, err := lcs.getNode("/html/body/div[1]/div/div[2]/div/div/div[1]/div/div[1]/div[1]/div/div[2]/div/div[2]/div/div")
	///html/body/div[1]/div/div[2]/div/div/div[1]/div/div[1]/div[1]/div/div[2]/div/div[2]/div/div//text()
	if err != nil {
		return nil, fmt.Errorf("Couldn't find example node")
	}
	fmt.Println("Scrapped exampleNodes; starting cleaning")
	lcp.Example = lcs.cleanScrapedString(lcs.concatNodeStr(exampleNodes))
	fmt.Println("Finished scrapping Examples")

	// Find related topics
	relatedTopicsNodes, err := lcs.getNode("/html/body/div[1]/div/div[2]/div/div/div[1]/div/div[1]/div[1]/div/div[2]/div/div[6]/div[2]/a/span/text()")
	if err != nil {
		return nil, fmt.Errorf("Couldn't find example node")
	}
	lcp.RelatedTopics = lcs.cleanScrapedString(lcs.concatNodeStr(relatedTopicsNodes))
	fmt.Println("Finished scrapping Related Topics")

	// Find difficulty
	difficultyNodes, err := lcs.getNode("/html/body/div[1]/div/div[2]/div/div/div[1]/div/div[1]/div[1]/div/div[2]/div/div[1]/div[2]/div/text()")
	if err != nil {
		return nil, fmt.Errorf("Couldn't find related topics node")
	}
	lcp.Difficulty = lcs.cleanScrapedString(difficultyNodes[0].Dump("", "", false))
	fmt.Println("Finished scrapping Difficulty")

	// Find related problems
	relatedProblemsNodes, err := lcs.getNode("/html/body/div[1]/div/div[2]/div/div/div[1]/div/div[1]/div[1]/div/div[2]/div/div[7]/div[2]/div/a/text()")
	if err != nil {
		return nil, fmt.Errorf("Couldn't find related problems node")
	}
	lcp.RelatedProblems = lcs.cleanScrapedString(lcs.concatNodeStr(relatedProblemsNodes))
	fmt.Println("Finished scrapping Related Topics")

	fmt.Println("Finished scrapping nodes")

	return &lcp, nil
}

func (lcs *LeetCodeScrapper) getNode(fullXPath string) ([]*cdp.Node, error) {
	var node []*cdp.Node
	fmt.Println("Getting node..")
	err := chromedp.Run(lcs.ctx,
		chromedp.Nodes(fullXPath, &node),
	)
	fmt.Println("Got node..")
	return node, err
}

func (lcs *LeetCodeScrapper) getText(fullXPath string) (*[]string, error) {
	jsText := jsGetText(fullXPath)
	var text *[]string
	chromedp.Evaluate(jsText, &text)
	return text, nil
}

func (lcs *LeetCodeScrapper) concatNodeStr(nodes []*cdp.Node) string {
	str := ""
	for _, n := range nodes {
		str += n.Dump(" ", "", false)
		fmt.Println(str)
	}
	return str
}

func (lcs *LeetCodeScrapper) cleanScrapedString(str string) string {
	cleanedStr := strings.ReplaceAll(str, "#text ", "")
	return strings.TrimSpace(strings.ReplaceAll(cleanedStr, `"`, ""))
}

func jsGetText(sel string) (js string) {
	const funcJS = `function getText(sel) {
				var text = [];
				var elements = document.body.querySelectorAll(sel);

				for(var i = 0; i < elements.length; i++) {
					var current = elements[i];
					if(current.children.length === 0 && current.textContent.replace(/ |\n/g,'') !== '') {
					// Check the element has no children && that it is not empty
						text.push(current.textContent + ',');
					}
				}
				return text
			 };`

	invokeFuncJS := `var a = getText('` + sel + `'); a;`
	return strings.Join([]string{funcJS, invokeFuncJS}, " ")
}
