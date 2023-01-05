package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"
)

type PageInfo struct {
	url      string
	content  string
	findings []string
}

func main() {
	timeStart := time.Now()

	pages := []PageInfo{
		{
			url:      "https://www.anais.digital/",
			content:  "",
			findings: []string{},
		},
		{
			url:      "https://www.anais.digital/cases/",
			content:  "",
			findings: []string{},
		},
		{
			url:      "https://www.anais.digital/expertises/",
			content:  "",
			findings: []string{},
		},
		{
			url:      "https://www.anais.digital/about-us/",
			content:  "",
			findings: []string{},
		},
		{
			url:      "https://www.anais.digital/expertise/ux-accessibility-academy/",
			content:  "",
			findings: []string{},
		},
		{
			url:      "https://www.anais.digital/trainings/accessibility/",
			content:  "",
			findings: []string{},
		},
		{
			url:      "https://www.anais.digital/expertise/digital-transformation/",
			content:  "",
			findings: []string{},
		},
		{
			url:      "https://www.anais.digital/expertise/digital-recrutement/",
			content:  "",
			findings: []string{},
		},
		{
			url:      "https://www.anais.digital/expertise/web-mobile-app/",
			content:  "",
			findings: []string{},
		},
		{
			url:      "https://www.anais.digital/expertise/user-experience/",
			content:  "",
			findings: []string{},
		},
		{
			url:      "https://www.anais.digital/jobs/",
			content:  "",
			findings: []string{},
		},
		{
			url:      "https://www.anais.digital/blog/",
			content:  "",
			findings: []string{},
		},
		{
			url:      "https://shows.acast.com/no-waste-podcast",
			content:  "",
			findings: []string{},
		},
	}
	pagesCount := len(pages)
	chPages := make(chan PageInfo)
	chResults := make(chan PageInfo)

	go processPages(chPages, chResults)
	go processPages(chPages, chResults)

	go func() {
		for pageIndex := 0; pageIndex < pagesCount; pageIndex++ {
			chPages <- pages[pageIndex]
		}
		close(chPages)
	}()

	for pageIndex := 0; pageIndex < pagesCount; pageIndex++ {
		page := <-chResults

		fmt.Printf("Page downloaded url: %v \n", page.url)
		fmt.Printf("Page content length downloaded: %v \n", len(page.content))
		fmt.Printf("Page findings: %v \n", page.findings)
		// pageSnippetSize := 750
		// fmt.Printf("Page content downloaded (only first %v bytes): %v \n", pageSnippetSize, string(pageContent[:pageSnippetSize]))
	}

	timeElapsed := time.Since(timeStart)
	fmt.Printf("Script took: %v \n", timeElapsed)

	fmt.Println()
}

func processPages(chPages <-chan PageInfo, chResults chan<- PageInfo) {
	for page := range chPages {
		pageContent, _ := getPage(page.url)

		page.content = string(pageContent)

		re := regexp.MustCompile(`foo.?`)
		findings := re.FindAllString(string(pageContent), -1)
		page.findings = findings

		chResults <- page
	}

}

func getPage(url string) ([]byte, error) {
	fmt.Println("trying to download page from url: " + url)
	res, err := http.Get(url)
	if err != nil {
		msg := fmt.Sprintf("error making http request: %s\n", err.Error())
		fmt.Printf(msg)
		return []byte(""), errors.New(msg)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	// fmt.Printf("Page content length downloaded: %v \n", len(body))
	// pageSnippetSize := 500
	// fmt.Printf("Page content downloaded (only first %v bytes): %v \n", pageSnippetSize, string(body[:pageSnippetSize]))

	return body, nil
}
