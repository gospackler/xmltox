package main

import (
	"encoding/base64"
	"fmt"
	"image/png"
	"io/ioutil"
)

// Use a reader for this to negate huge slice created.
func decodeBase64(base64Data []byte) (dst []byte) {

	dst = make([]byte, 1000000)
	n, err := base64.StdEncoding.Decode(dst, base64Data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Value of n = ", n)
	return
}

func main() {
	base64Data, _ := ioutil.ReadFile("base64encode")
	pngData := decodeBase64(base64Data)
	ioutil.WriteFile("goabc.png", pngData, 0644)
}
