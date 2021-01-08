package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

var paths []string

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "The url that you want to build a sitemap for")
	flag.Parse()

	P := search(*urlFlag, "/")
	P_unique := unique(P)
	for n, p := range P_unique {
		fmt.Printf("%d -> %s%s\n", n, *urlFlag, p)
	}

}

func search(url string, path string) []string {

	full_link := url + path
	resp_body := getUrl(full_link)

	links := getLinks(resp_body)
	unique_links := unique(links)
	internal_links := getInternalLinks(unique_links)

	for _, path := range internal_links {
		paths = append(paths, path)
		search(url, path)
	}

	return paths

}

func getUrl(url string) string {
	resp, err := http.Get(url)
	check(err)

	resp_body, err := ioutil.ReadAll(resp.Body)
	check(err)

	return string(resp_body)
}

func getLinks(resp_body string) []string {
	body := strings.NewReader(resp_body)
	var links []string
	z := html.NewTokenizer(body)

	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}
				}
			}
		}
	}
	return links
}

func getInternalLinks(links []string) []string {
	var internalLinks []string
	for _, link := range links {
		if link == "/" {
			continue
		} else if strings.HasPrefix(link, "/") {
			internalLinks = append(internalLinks, link)
		}
	}
	return internalLinks
}

func unique(input []string) []string {
	keys := make(map[string]bool)
	output := []string{}
	for _, entry := range input {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			output = append(output, entry)
		}
	}
	return output
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
