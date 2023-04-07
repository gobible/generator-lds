package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

func getLOF() {

	book := b.AddBook("Lectures on Faith")
	book.Notes = getContent("https://www.centerplace.org/hs/lf/lec-hist.htm")

	// 1-144, with 108 having 108a as well
	// https://www.centerplace.org/hs/dc/section002.htm

	for i := 1; i <= 7; i++ {
		fmt.Println("Processing LOF section", i)
		url := fmt.Sprintf("https://www.centerplace.org/hs/lf/lecture%03d.htm", i)
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

		converter := md.NewConverter("", true, nil)
		markdown := converter.Convert(doc.Find("div.content"))

		section := book.GetChapterOrCreate(i)
		section.Name = "Section " + strconv.Itoa(i)
		section.Content = markdown

	}
}
