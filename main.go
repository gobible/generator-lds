/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	gobible "github.com/applehat/gobible"
	bible "github.com/applehat/gobible/bible"
)

var b *bible.Bible

func main() {

	b = gobible.LoadInternal("KJV")

	b.Version.Name = "The Inspired Version (Joseph Smith Translation)"
	b.Version.Abbrev = "IV"
	b.Extra.Foreword = getContent("http://www.centerplace.org/hs/iv/foreword.htm")
	b.Extra.Preface = getContent("http://www.centerplace.org/hs/iv/preface.htm")

	// First lets get the perversion of holy scripture
	for _, book := range bible.BooksTable {
		fmt.Println("Processing", book.Name, "")
		if book.Name == "Song of Solomon" {
			b.GetBook(book.Name).Notes = "Song of Solomon was not modified by Joseph Smith in the Inspired Version as far as we know. The included text is from the King James Version."
			continue
		}
		b.GetBook(book.Name).Notes = "This is an corruption of the original text of " + book.Name + ", by Joseph Smith. No Christian should use this for anything but cult study purposes."
		getBibleBook(book.Name)
	}

	// Now lets get the Book of Mormon
	abominations := []string{
		"1 Nephi",
		"2 Nephi",
		"Jacob",
		"Enos",
		"Jarom",
		"Omni",
		"Words of Mormon",
		"Mosiah",
		"Alma",
		"Helaman",
		"3 Nephi",
		"4 Nephi",
		"Mormon",
		"Ether",
		"Moroni",
	}

	for _, book := range abominations {
		fmt.Println("Processing", book, "")
		getBOMBook(book)
	}

	// Get the Doctrine and Covenants
	getDC()

	// Get Lectures on Faith
	getLOF()

	file, _ := json.MarshalIndent(b, "", " ")
	_ = os.WriteFile("lds-scriptures.json", file, 0644)
}

func getContent(url string) string {
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

	doc.Find("button").Remove()

	converter := md.NewConverter("", true, nil)
	markdown := converter.Convert(doc.Find("body"))
	return markdown
}
