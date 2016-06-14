package xmltox

// #include "xmltox.h"
// #cgo LDFLAGS: -lwkhtmltox -lxml2 -lxslt
import "C"

import (
	"errors"
	"io/ioutil"
)

func init() {
	C.wkpdfInit()
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

func GetPNG(xmlContent []byte, xslContent []byte, uidFileName string) (png []byte, err error) {

	// Add a prefix path to the uidFileName
	xmlFileName := uidFileName + ".xml"
	err = ioutil.WriteFile(xmlFileName, xmlContent, 0644)
	if err != nil {
		return nil, err
	}

	xslFileName := uidFileName + ".xsl"
	err = ioutil.WriteFile(xslFileName, xslContent, 0644)
	if err != nil {
		return nil, err
	}

	selDriver, err := initSelDriver("input/e1_inpatient.xml", "input/CED.XSL")
	if err != nil {
		return nil, err
	}
	png, err = selDriver.TakeScreenshot()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return
}
