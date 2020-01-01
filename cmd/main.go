package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
)

type LeetCodeProblem struct {
	Name          string
	Number        string
	Description   string
	Example       string
	Difficulty    string
	Url           string
	RelatedTopics string
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

func toSlug(str string) string {
	return strings.ReplaceAll(strings.ToLower(str), " ", "-")
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
	URL      string
	document *goquery.Document
}

func NewLeetCodeScrapper(url string) *LeetCodeScrapper {
	lcs := LeetCodeScrapper{
		URL: url,
	}
	// lcs.getHTML()
	return &lcs
}

func (lcs *LeetCodeScrapper) getHTML() {

}

func (lcs *LeetCodeScrapper) ScrapeData() (*LeetCodeProblem, error) {
	lcp := LeetCodeProblem{}
	fmt.Printf("%+v", lcs.document)

	opts := append(chromedp.DefaultExecAllocatorOptions[:]) // chromedp.Flag("headless", false),
	// chromedp.Flag("disable-gpu", false),
	// chromedp.Flag("enable-automation", false),
	// chromedp.Flag("disable-extensions", false),

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create context
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	if err := chromedp.Run(ctx,
		chromedp.Emulate(device.IPhone7landscape),
		chromedp.Navigate(lcs.URL),
		chromedp.Sleep(2000*time.Millisecond),

		// chromedp.WaitVisible(`#thing`),
	); err != nil {
		return nil, fmt.Errorf("Couldn't run chrome browser")
	}

	// Find name & number
	var nameNode []*cdp.Node
	if err := chromedp.Run(ctx,
		chromedp.Nodes("//h3/text()", &nameNode),
	); err != nil {
		return nil, fmt.Errorf("Couldn't find")
	}
	// Name
	lcp.Name = nameNode[2].Dump("", "", false)
	fmt.Println(lcp.Name)

	// Number
	lcp.Number = nameNode[0].Dump("", "", false)
	fmt.Println(lcp.Number)

	fmt.Println("Document tree:")
	fmt.Print(nameNode[0].Value)
	fmt.Println("Finished printing nodes")

	return &lcp, nil
}
