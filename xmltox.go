package xmltox

// #include "xmltox.h"
// #cgo LDFLAGS: -lwkhtmltox -lxml2 -lxslt
import "C"

import (
	"errors"
	"io/ioutil"
)

func GetHTML(xmlContent []byte, xslContent []byte) (html []byte, err error) {
	cxml := C.CString(string(xmlContent))
	cxsl := C.CString(string(xslContent))

	html = []byte(C.GoString(C.GetHTML(cxml, cxsl)))
	if html != nil {
		errors.New("html not obtained from GetHTML")
	}
	return
}

func GetPDF(xmlContent []byte, xslContent []byte) (pdf []byte, err error) {
	cxml := C.CString(string(xmlContent))
	cxsl := C.CString(string(xslContent))

	fileName := C.GoString(C.GetPDFFile(cxml, cxsl))
	if fileName == "" {
		return nil, errors.New("PDF not obtained from GetPDFData")
	}

	pdf, err = ioutil.ReadFile(fileName)
	if err != nil {
		return nil, errors.New("Error reading created pdf")
	}
	return
}

func GetPNG(xmlContent []byte, xslContent []byte) (png []byte, err error) {
	return
}
