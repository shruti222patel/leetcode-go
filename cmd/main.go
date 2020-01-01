package leetcodegen

import (
	"flag"

	"./templater"
)

func main() {
	// Get leetcode url to parse from cli
	url := flag.String("url", "https://leetcode.com/problems/add-two-numbers/", "Url of the leetcode problem")
	flag.Parse()

	if url == nil {
		panic("Couldn't find url from cli args")
	}

	err := templater.GenerateTemplates(*url)
	if err != nil {
		panic(err)
	}

}
