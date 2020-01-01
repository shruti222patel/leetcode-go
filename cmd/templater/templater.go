package templater

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"../scrapper"
)

func GenerateTemplates(url string) error {
	// lc := getLeetCodeProblem("dummy url")

	lcs := scrapper.NewLeetCodeScrapper(url)
	lcp, err := lcs.ScrapeData()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", lcp)

	t := template.Must(template.ParseGlob("**/*.tmpl"))

	// Create readme.md
	err = t.ExecuteTemplate(os.Stdout, "readme.tmpl", lcp)
	if err != nil {
		return err
	}

	// Create main.go
	err = t.ExecuteTemplate(os.Stdout, "main.tmpl", lcp)
	if err != nil {
		return err
	}

	// Create main_test.go
	err = t.ExecuteTemplate(os.Stdout, "main_test.tmpl", lcp)
	if err != nil {
		return err
	}

	return nil
}

func toSlug(str string) string {
	return strings.ReplaceAll(strings.ToLower(str), " ", "-")
}
