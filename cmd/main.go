package main

import (
	"flag"
	"fmt"

	"./templater"
)

func main() {
	fmt.Println("Starting main fn..")
	// Get leetcode url to parse from cli
	url := flag.String("url", "", "Url of the leetcode problem")
	flag.Parse()

	if url == nil || *url == "" {
		panic("Couldn't find url from cli args")
	}

	fmt.Println("Starting template generation..")
	err := templater.GenerateTemplates(*url)
	if err != nil {
		panic(err)
	}

}
