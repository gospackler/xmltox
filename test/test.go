package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/jung-kurt/gofpdf"
)

// Use a reader for this to negage slice created.
// Using the make is suicide with a number passed.
func decodeBase64(base64Data []byte) (dst []byte) {

	dst = make([]byte, 1000000)
	n, err := base64.StdEncoding.Decode(dst, base64Data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Value of n = ", n)
	return
}

// Return the pdf data in bytes.
// Ideally put the png into the pdf and generate it.
func getPdf(pngData []byte) ([]byte, error) {
	pngReader := bytes.NewReader(pngData)
	pdf := gofpdf.New("P", "mm", "A4", "")
	options := gofpdf.ImageOptions{
		ReadDpi:   false,
		ImageType: "PNG",
	}
	imageName := "info"
	pdf.RegisterImageOptionsReader(imageName, options, pngReader)
	pdf.ImageOptions(imageName, 0, 0, 10, 10, true, options, 1, "")

	pdf.OutputFileAndClose("goabc.pdf")
	return nil, nil
}

func main() {
	base64Data, _ := ioutil.ReadFile("base64encode")
	pngData := decodeBase64(base64Data)
	ioutil.WriteFile("goabc.png", pngData, 0644)
	getPdf(pngData)
}
