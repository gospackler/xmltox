package xmltox

// #include "xmltox.h"
// #cgo LDFLAGS: -lwkhtmltox -lxml2 -lxslt
import "C"

import (
	"errors"
	"io/ioutil"
	//	"fmt"
	//	"unsafe"
	//	"os"
)

func GetHTML(xmlContent []byte, xslContent []byte) (html []byte, err error) {

	cxml := C.CString(string(xmlContent))
	cxsl := C.CString(string(xslContent))

	uidFileName := "uidfilename"
	status := C.InitStatus(C.CString(uidFileName), cxml, cxsl)
	success := C.GetHTML(status)
	if !success {
		return nil, errors.New("Error generating html")
	}

	fileName := uidFileName + ".html"
	html, err = ioutil.ReadFile(fileName)
	if err != nil {
		return nil, errors.New("Error reading created html" + err.Error())
	}
	defer C.FinishStatus(status)
	return
}

func GetPDF(xmlContent []byte, xslContent []byte, uidFileName string) (pdf []byte, err error) {

	cxml := C.CString(string(xmlContent))
	cxsl := C.CString(string(xslContent))

	status := C.InitStatus(C.CString(uidFileName), cxml, cxsl)
	success := C.GetHTML(status)
	if !success {
		return nil, errors.New("Error generating html")
	}

	success = C.GetPDFFile(status)
	if !success {
		return nil, errors.New("PDF not obtained from GetPDFData")
	}

	fileName := uidFileName + ".pdf"
	pdf, err = ioutil.ReadFile(fileName)
	if err != nil {
		return nil, errors.New("Error reading created pdf" + err.Error())
	}

	defer C.FinishStatus(status)
	return
}

func GetPNG(xmlContent []byte, xslContent []byte) (png []byte, err error) {

	cxml := C.CString(string(xmlContent))
	cxsl := C.CString(string(xslContent))

	uidFileName := "uidfilename"
	status := C.InitStatus(C.CString(uidFileName), cxml, cxsl)
	success := C.GetHTML(status)
	if !success {
		return nil, errors.New("Error generating html")
	}

	fileName := uidFileName + ".html"
	cData := C.CString("")
	//	len := C.GenPNG(status)
	//	cData := C.GetPNG(status)

	C.WkpngCreate(C.CString(fileName), &cData)
	//len := C.WkpngCreate(C.CString(fileName), &cData)
	//	png = C.GoBytes(unsafe.Pointer(cData), len)

	//	fmt.Println("Am I being seen?")
	//	C.FinishStatus(status)

	return
}
