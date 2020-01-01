package main

import (
	"fmt"
	"strings"

	"./scrapper"
)

func main() {
	// // Get leetcode url to parse from cli
	// url := flag.String("url", "", "Url of the leetcode problem")
	// flag.Parse()

	lcs := scrapper.NewLeetCodeScrapper("https://leetcode.com/problems/add-two-numbers/")

	lcp, err := lcs.ScrapeData()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", lcp)
}

func getLeetCodeProblem(url string) *scrapper.LeetCodeProblem {
	lc := scrapper.LeetCodeProblem{
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

func toSlug(str string) string {
	return strings.ReplaceAll(strings.ToLower(str), " ", "-")
}
