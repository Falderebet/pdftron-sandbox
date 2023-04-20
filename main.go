package main

import (
	"fmt"

	. "github.com/pdftron/pdftron-go/v2"
)

func main() {
	PDFNetInitialize("demo:1681479144517:7df72d550300000000cd5af736e910cc63e454cadb25e0d06c2b7147ee")

	filename := "PDFTRON-test-pdf.pdf"

	// open document from the filesystem
	doc := NewPDFDoc(filename)

	fmt.Println(doc.GetFileName())

}
