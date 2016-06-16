package main

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io"
	"io/ioutil"

	"github.com/jung-kurt/gofpdf"
	"github.com/oliamb/cutter"
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

func cropImage(bigImage image.Image, startX, startY int, height int) (io.Reader, error) {
	width := bigImage.Bounds().Dx()

	croppedImg, err := cutter.Crop(bigImage, cutter.Config{
		Width:   width,
		Height:  height,
		Anchor:  image.Point{startX, startY},
		Options: cutter.Copy,
	})
	if err != nil {
		return nil, errors.New("Crop Error : " + err.Error())
	}
	var tempByte []byte
	buffer := bytes.NewBuffer(tempByte)
	err = png.Encode(buffer, croppedImg)
	if err != nil {
		return nil, errors.New("Image Encode Error : " + err.Error())
	}

	reader := bytes.NewReader(buffer.Bytes())
	fmt.Println("Reader length = ", reader.Len())
	return reader, nil
}

func scaleBigImage(bigImage image.Image, id int, total int) (reader io.Reader, err error) {

	if id > total || id < 0 {
		return nil, errors.New("id should be less than total positive i.e 0 < id < total")
	}

	height := bigImage.Bounds().Dy()
	startY := (height / total) * id
	reader, err = cropImage(bigImage, 0, startY, height/total)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return
}

func addA4PdfPage(pdf *gofpdf.Fpdf, pageName string, reader io.Reader) {
	options := gofpdf.ImageOptions{
		ReadDpi:   false,
		ImageType: "PNG",
	}
	widthA4 := 11.7
	heightA4 := 8.27
	imageInfo := pdf.RegisterImageOptionsReader(pageName, options, reader)
	dpi := 72.0 // as per the docs of the api
	imageInfo.SetDpi(dpi)
	pdf.ImageOptions(pageName, 0, 0, dpi*widthA4, dpi*heightA4, true, options, 0, "")
}

// Return the pdf data in bytes.
// Ideally put the png into the pdf and generate it.
func getPdf(pngData []byte) ([]byte, error) {
	pdf := gofpdf.New("L", "pt", "A4", "")

	pngReader := bytes.NewReader(pngData)
	bigImage, err := png.Decode(pngReader)
	if err != nil {
		return nil, errors.New("Error PNG Decode : " + err.Error())
	}

	reader1, err := scaleBigImage(bigImage, 0, 3)
	if err != nil {
		return nil, err
	}

	reader2, err := scaleBigImage(bigImage, 1, 3)
	if err != nil {
		return nil, err
	}

	reader3, err := scaleBigImage(bigImage, 2, 3)
	if err != nil {
		return nil, err
	}
	addA4PdfPage(pdf, "page1", reader1)
	addA4PdfPage(pdf, "page2", reader2)
	addA4PdfPage(pdf, "page3", reader3)

	pdf.OutputFileAndClose("goabc.pdf")
	fmt.Println("Processing status", pdf.Ok())
	return nil, nil
}

func main() {
	base64Data, _ := ioutil.ReadFile("base64encode")
	pngData := decodeBase64(base64Data)
	ioutil.WriteFile("goabc.png", pngData, 0644)
	_, err := getPdf(pngData)
	if err != nil {
		panic(err)
	}
}
