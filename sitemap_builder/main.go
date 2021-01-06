package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "The url that you want to build a sitemap for")
	flag.Parse()

	resp, err := http.Get(*urlFlag)
	check(err)

	for _, v := range getLinks(resp.Body) {
		fmt.Println(v)
	}

}

func getLinks(body io.Reader) []string {
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

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
