package xmltox

// #include "test.h"
import "C"

import (
	"fmt"
)

func GetHTML(xmlContent []byte) (html []byte, err error) {

	fmt.Println("sum =", C.Sum(5, 10))
	return
}

func GetPDF(xmlContent []byte) (pdf []byte, err error) {
	return
}

func GetPNG(xmlContent []byte) (png []byte, err error) {
	return
}
