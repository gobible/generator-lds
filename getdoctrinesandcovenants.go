package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

func getDC() {

	book := b.AddBook("Doctrine and Covenants")
	book.Notes = getContent("http://www.centerplace.org/hs/dc/preface.htm")

	// 1-144, with 108 having 108a as well
	// https://www.centerplace.org/hs/dc/section002.htm

	for i := 1; i <= 144; i++ {
		fmt.Println("Processing section", i, "of 144")
		url := fmt.Sprintf("https://www.centerplace.org/hs/dc/section%03d.htm", i)
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

		section := book.GetChapterOrCreate(i)
		section.Name = "Section " + strconv.Itoa(i)

		converter := md.NewConverter("", true, nil)
		markdown := converter.Convert(doc.Find("div.content"))

		section.Name = "Section " + strconv.Itoa(i)
		section.Content = markdown
	}
	fmt.Println("Processing section 108a")
	url := fmt.Sprintf("https://www.centerplace.org/hs/dc/section108a.htm")
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

	section := book.GetChapterOrCreate(145)
	section.Name = "Section 108a"

	converter := md.NewConverter("", true, nil)
	markdown := converter.Convert(doc.Find("div.content"))

	section.Name = "Section 108a"
	section.Content = markdown

}
