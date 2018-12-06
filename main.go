package main

import (
	"strings"
	// "fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)
// YC check for yc headlines, news, ask
func YC() {
	// Request the HTML page.
	ws := []string{"http://news.ycombinator.com/news", "http://news.ycombinator.com/newest"}
	for _, v := range ws {
		res, err := http.Get(v)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Find the review items
		doc.Find("a.storylink").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			// println(s.Text())
			// println("yes")
			t := []string{"btc", "bitcoin", "Yale", "Work", "Bitcoin", "BTC"}
			for _, u := range t {
				if strings.Contains(s.Text(), u) {
					println(s.Text())
					if ht, e := s.Html(); e != nil {
						println(e)
					} else {
						println(ht)
					}
				}
			}
			// band:=strings.Contains(s.Text(), "Yale")
			// fmt.Printf("Review %d:%s, %s - %s\n", i, "", band, title)
		})
	}
}

func main() {
	YC()
}
