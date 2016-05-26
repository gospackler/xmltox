package xmltox

import (
	"testing"

	"io/ioutil"
)

const XmlLocation = "input/e1_inpatient.xml"
const XslLocation = "input/CED.XSL"

func TestGetHTML(t *testing.T) {

	b, err := ioutil.ReadFile(XmlLocation)
	if err != nil {
		t.Errorf("Error reading xml file", XmlLocation)
	}
	// for now TODO
	html, err := GetHTML(b)
	if err != nil {
		t.Errorf("Error getting html")
	}
	t.Log(html)
}

func TestGetPDF(t *testing.T) {

	b, err := ioutil.ReadFile(XmlLocation)
	if err != nil {
		t.Errorf("Error reading xml file")
	}
	pdf, err := GetPDF(b)
	if err != nil {
		t.Errorf("Error getting pdf")
	}
	t.Log(pdf)
}

func TestGetPNG(t *testing.T) {

	b, err := ioutil.ReadFile(XmlLocation)
	if err != nil {
		t.Errorf("Error reading xml file")
	}
	png, err := GetPNG(b)
	if err != nil {
		t.Errorf("Error getting png")
	}
	t.Log(png)
}
