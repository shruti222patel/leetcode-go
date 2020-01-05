package templater

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"

	"../scrapper"
)

func GenerateTemplates(url string) error {
	// lc := getLeetCodeProblem("dummy url")

	lcs := scrapper.NewLeetCodeScrapper(url, true)
	lcp, err := lcs.ScrapeData()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", lcp)

	t := template.Must(template.ParseGlob("**/*.tmpl"))

	directoryName, err := makeDir(lcp.Name, lcp.Number)
	if err != nil {
		return err
	}

	// Create readme.md
	err = createFile(t, lcp, directoryName, "readme.md.tmpl")
	// err = t.ExecuteTemplate(os.Stdout, "readme.tmpl", lcp)
	if err != nil {
		return err
	}

	// Create main.go
	err = createFile(t, lcp, directoryName, "main.go.tmpl")
	// err = t.ExecuteTemplate(os.Stdout, "main.tmpl", lcp)
	if err != nil {
		return err
	}

	// Create main_test.go
	err = createFile(t, lcp, directoryName, "main_test.go.tmpl")
	// err = t.ExecuteTemplate(os.Stdout, "main_test.tmpl", lcp)
	if err != nil {
		return err
	}

	return nil
}

func makeDir(prbName, prbNumber string) (string, error) {
	// func getDirectoryName(url string, string) string {
	// splitStr := strings.Split(url, "/")
	// lastStr := strings.TrimSpace(splitStr[len(splitStr)-1])
	// if lastStr == "" {
	// 	return strings.TrimSpace(splitStr[len(splitStr)-2])
	// }
	// return lastStr

	dirName := prbNumber + "." + toSlug(prbName)

	err := os.Mkdir("../pkg/"+dirName, 0755)
	if err != nil {
		return "", err
	}
	return dirName, nil
}

func createFile(t *template.Template, lcp *scrapper.LeetCodeProblem, directoryName, templateName string) error {
	fileName := strings.ReplaceAll(templateName, ".tmpl", "")
	w, err := getWriter("../pkg/" + directoryName + "/" + fileName)
	if err != nil {
		return err
	}
	return t.ExecuteTemplate(w, templateName, lcp)
}

func getWriter(filePath string) (*bufio.Writer, error) {
	f, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}

	return bufio.NewWriter(f), nil
}

func toSlug(str string) string {
	return strings.ReplaceAll(strings.ToLower(str), " ", "-")
}
