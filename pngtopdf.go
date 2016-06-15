package xmltox

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"image/png"
	"io"

	"github.com/jung-kurt/gofpdf"
	"github.com/oliamb/cutter"
)

// Use a reader for this to negage slice created.
// Using the make is suicide with a number passed.
func decodeBase64(base64Data string) (dst []byte, err error) {

	dst, err = base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return nil, err
	}
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
	pdf.RegisterImageOptionsReader(pageName, options, reader)
	width, height := pdf.GetPageSize()
	pdf.ImageOptions(pageName, 0, 0, width, height, true, options, 0, "")
}

// Return the pdf data in bytes.
// Ideally put the png into the pdf and generate it.
func getPDF(pngData []byte, numOfPages int) ([]byte, error) {
	pdf := gofpdf.New("P", "pt", "A4", "")
	pdf.SetMargins(0, 0, 0)
	pngReader := bytes.NewReader(pngData)
	bigImage, err := png.Decode(pngReader)
	if err != nil {
		return nil, errors.New("Error PNG Decode : " + err.Error())
	}

	for pageNo := 0; pageNo < numOfPages; pageNo++ {
		pageReader, err := scaleBigImage(bigImage, 0, numOfPages)
		if err != nil {
			return nil, err
		}
		addA4PdfPage(pdf, "page"+string(pageNo), pageReader)
	}
	var temp []byte
	buffer := bytes.NewBuffer(temp)
	pdf.Output(buffer)
	return buffer.Bytes(), nil
}

/*
func main() {
	base64Data, _ := ioutil.ReadFile("base64encode")
	pngData, _ := decodeBase64(base64Data)
	ioutil.WriteFile("goabc.png", pngData, 0644)
	_, err := getPdf(pngData)
	if err != nil {
		panic(err)
	}
}
*/
