package xmltox

// #include "xmltox.h"
// #cgo LDFLAGS: -lwkhtmltox -lxml2 -lxslt
import "C"

import (
	"errors"
	"io/ioutil"

	"unsafe"

	"fmt"

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
	C.FinishStatus(status)
	return
}

func GetPDF(xmlContent []byte, xslContent []byte) (pdf []byte, err error) {

	cxml := C.CString(string(xmlContent))
	cxsl := C.CString(string(xslContent))

	uidFileName := "uidfilename"
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

	C.FinishStatus(status)
	return
}

func GetPNG(xmlContent []byte, xslContent []byte) (png []byte, err error) {

	cxml := C.CString(string(xmlContent))
	cxsl := C.CString(string(xslContent))

	status := C.InitStatus(C.CString("uidfilename"), cxml, cxsl)
	success := C.GetHTML(status)
	if !success {
		return nil, errors.New("Error generating html")
	}
	len := C.GenPNG(status)
	cData := C.GetPNG(status)

	png = C.GoBytes(unsafe.Pointer(cData), len)

	zeros := 0
	for i := 0; i < int(len); i++ {
		if png[i] == 0 {
			zeros = zeros + 1
		}
	}

	fmt.Println("zero count = ", zeros)
	// No need to check for retyrn as it never realy returns false.
	//	C.FinishStatus(status)

	return
}
