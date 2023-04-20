package main

import (
	"fmt"

	. "github.com/pdftron/pdftron-go/v2"
)

const filename = "PDFTRON-test-pdf.pdf"

func main() {

	PDFNetInitialize("demo:1681479144517:7df72d550300000000cd5af736e910cc63e454cadb25e0d06c2b7147ee")

	// open document from the filesystem
	doc := NewPDFDoc(filename)

	// Printing the filename
	fmt.Println(doc.GetFileName())

	//Read text from a file example from - https://docs.apryse.com/documentation/linux/guides/features/extraction/text-extract/
	readTextFromFile()

}

func readTextFromFile() {
	doc := NewPDFDoc(filename)
	page := doc.GetPage(1)

	txt := NewTextExtractor()
	txt.Begin(page) // Readspage

	// Extract words one by one
	word := NewWord()
	line := txt.GetFirstLine()
	for line.IsValid() {
		word = line.GetFirstWord()
		for word.IsValid() {
			// to get words in string you need to call .GetString() on a "word".
			fmt.Println("This is the current word: ", word.GetString())
			word = word.GetNextWord()
		}
		// To see new lines not needed
		fmt.Println()
		line = line.GetNextLine()
	}

}
