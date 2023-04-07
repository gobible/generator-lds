package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getBibleBook(bookName string) {
	fmt.Println("Getting book", bookName)
	url := fmt.Sprintf("http://www.centerplace.org/hs/iv/%s.htm", strings.ToLower(strings.ReplaceAll(bookName, " ", "")))
	res, err := http.Get(url)
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

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("name")
		if name == "" {
			return
		}

		first := name[0:1]
		rest := name[1:]
		// get first char of name

		// find first P tag

		if first == "c" {
			// set book title
			cn, _ := strconv.Atoi(rest)

			chapterTitle := s.Next().Text()

			//fmt.Println("Setting chapter title for chapter", cn, ":", chapterTitle)
			b.GetBook(bookName).GetChapter(cn).SetTitle(chapterTitle)
			//fmt.Println("Stored: ", b.GetBook(book.Name).GetChapter(cn).Title)

		}
		if first == "v" {
			fmt.Println("Processing ", bookName, " ", rest)
			split := strings.Split(rest, ".")
			cn, _ := strconv.Atoi(split[0])
			vn, _ := strconv.Atoi(split[1])
			text := strings.ReplaceAll(s.Text(), split[0]+":"+split[1]+" ", "")

			//fmt.Println("Verse ", rest, ":", text)
			// set verse text

			b.GetBook(bookName).GetChapter(cn).SetVerse(vn, text)
		}
	})
}
