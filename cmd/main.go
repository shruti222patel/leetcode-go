package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
)

type LeetCodeProblem struct {
	Name            string
	Number          string
	Description     string
	Example         string
	Difficulty      string
	Url             string
	RelatedTopics   string
	RelatedProblems string
}

func main() {
	// // Get leetcode url to parse from cli
	// url := flag.String("url", "", "Url of the leetcode problem")
	// flag.Parse()

	lcs := NewLeetCodeScrapper("https://leetcode.com/problems/add-two-numbers/")

	lcp, err := lcs.ScrapeData()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", lcp)
}

func generateTemplates() {
	lc := getLeetCodeProblem("dummy url")

	t := template.Must(template.ParseGlob("*.tmpl"))

	err := t.ExecuteTemplate(os.Stdout, "main.tmpl", lc)
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(os.Stdout, "readme.tmpl", lc)
	err = t.ExecuteTemplate(os.Stdout, "main_test.tmpl", lc)
}

func getLeetCodeProblem(url string) *LeetCodeProblem {
	lc := LeetCodeProblem{
		Name:          "Test",
		Number:        "0003",
		Description:   "test",
		Example:       "test",
		Difficulty:    "test",
		Url:           url,
		RelatedTopics: "test",
	}
	return &lc
}

type LeetCodeScrapper struct {
	URL string
	ctx context.Context
}

func NewLeetCodeScrapper(url string) *LeetCodeScrapper {
	lcs := LeetCodeScrapper{
		URL: url,
	}

	return &lcs
}

func (lcs *LeetCodeScrapper) ScrapeData() (*LeetCodeProblem, error) {
	lcp := LeetCodeProblem{}

	opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", false))
	// chromedp.Flag("disable-gpu", false),
	// chromedp.Flag("enable-automation", false),
	// chromedp.Flag("disable-extensions", false),

	allocCtx, cancelAlloc := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancelAlloc()

	// create context
	ctx, cancelCtx := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancelCtx()

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

	// Find description
	descNodes, _ := lcs.getNode("/html/body/div[1]/div/div[2]/div/div/div[1]/div/div[1]/div[1]/div/div[2]/div/div[2]/div/p/text()")
	if err != nil {
		return nil, fmt.Errorf("Couldn't find descp node")
	}
	lcp.Description = lcs.cleanScrapedString(lcs.concatNodeStr(descNodes))

	// Find explaination
	exampleNodes, err := lcs.getNode("/html/body/div[1]/div/div[2]/div/div/div[1]/div/div[1]/div[1]/div/div[2]/div/div[2]/div/pre/text()")
	if err != nil {
		return nil, fmt.Errorf("Couldn't find example node")
	}
	lcp.Example = lcs.cleanScrapedString(exampleNodes[0].Dump("", "", false))

	// Find related topics
	relatedTopicsNodes, err := lcs.getNode("/html/body/div[1]/div/div[2]/div/div/div[1]/div/div[1]/div[1]/div/div[2]/div/div[6]/div[2]/a/span/text()")
	if err != nil {
		return nil, fmt.Errorf("Couldn't find example node")
	}
	lcp.RelatedTopics = lcs.cleanScrapedString(lcs.concatNodeStr(relatedTopicsNodes))

	// Find difficulty
	difficultyNodes, err := lcs.getNode("/html/body/div[1]/div/div[2]/div/div/div[1]/div/div[1]/div[1]/div/div[2]/div/div[1]/div[2]/div/text()")
	if err != nil {
		return nil, fmt.Errorf("Couldn't find related topics node")
	}
	lcp.Difficulty = lcs.cleanScrapedString(difficultyNodes[0].Dump("", "", false))

	// Find related problems
	relatedProblemsNodes, err := lcs.getNode("/html/body/div[1]/div/div[2]/div/div/div[1]/div/div[1]/div[1]/div/div[2]/div/div[7]/div[2]/div/a/text()")
	if err != nil {
		return nil, fmt.Errorf("Couldn't find related problems node")
	}
	// TODO loop over all related problems
	lcp.RelatedProblems = lcs.cleanScrapedString(lcs.concatNodeStr(relatedProblemsNodes))

	fmt.Println("Finished printing nodes")

	return &lcp, nil
}

func (lcs *LeetCodeScrapper) getNode(fullXPath string) ([]*cdp.Node, error) {
	var node []*cdp.Node
	err := chromedp.Run(lcs.ctx,
		chromedp.Nodes(fullXPath, &node),
	)
	return node, err
}

func (lcs *LeetCodeScrapper) concatNodeStr(nodes []*cdp.Node) string {
	str := ""
	for _, n := range nodes {
		str += n.Dump(" ", "", false)
	}
	return str
}

func (lcs *LeetCodeScrapper) cleanScrapedString(str string) string {
	cleanedStr := strings.ReplaceAll(str, "#text ", "")
	return strings.ReplaceAll(cleanedStr, `"`, "")
}

func toSlug(str string) string {
	return strings.ReplaceAll(strings.ToLower(str), " ", "-")
}
