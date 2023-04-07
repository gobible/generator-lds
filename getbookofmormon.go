package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getBOMBook(bookName string) {
	book := b.AddBook(bookName)

	url := fmt.Sprintf("http://www.centerplace.org/hs/bm/%s.htm", strings.ToLower(strings.ReplaceAll(bookName, " ", "")))
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

	book.Notes = "NOTE: The book " + book.Name + " is a false testement attributed to Jesus Christ as part of the Book of Mormon. It is not part of the Bible."

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("name")
		if name == "" {
			return
		}
		first := name[0:1]
		rest := name[1:]
		if first == "c" {
			// set book title
			cn, _ := strconv.Atoi(rest)
			chapterTitle := s.Next().Text()
			book.Notes = book.Notes + "\n\n" + s.Prev().Text()
			book.GetChapterOrCreate(cn).Name = bookName + " " + rest
			book.GetChapterOrCreate(cn).SetTitle(chapterTitle)
		}
		if first == "v" {
			fmt.Println("Processing ", book.Name, " ", rest)
			split := strings.Split(rest, ".")
			cn, _ := strconv.Atoi(split[0])
			vn, _ := strconv.Atoi(split[1])
			text := strings.ReplaceAll(s.Text(), split[0]+":"+split[1]+" ", "")
			b.GetBook(book.Name).GetChapterOrCreate(cn).SetVerse(vn, text)
		}
	})
}
