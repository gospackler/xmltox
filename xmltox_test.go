package xmltox

import (
	"testing"

	"io/ioutil"
)

const XmlLocation = "input/e1_inpatient.xml"
const XslLocation = "input/CED.XSL"

func getXslXml(t *testing.T) ([]byte, []byte) {

	xml, err := ioutil.ReadFile(XmlLocation)
	if err != nil {
		t.Errorf("Error reading xml file", XmlLocation)
	}

	xsl, err := ioutil.ReadFile(XslLocation)
	if err != nil {
		t.Errorf("Error reading xml file", XmlLocation)
	}
	return xml, xsl
	// for now TODO
}

func TestGetPDF(t *testing.T) {

	xml, xsl := getXslXml(t)
	pdf, err := GetPDF(xml, xsl, "uid")
	if err != nil {
		t.Errorf("Error getting pdf")
	}
	t.Log("Received pdf bytes of length ", len(pdf))
	err = ioutil.WriteFile("testdatapdf.pdf", pdf, 0644)
	if err != nil {
		t.Errorf("Error writing pdf")
	}

	xml, xsl = getXslXml(t)
	pdf, err = GetPDF(xml, xsl, "uid1")
	if err != nil {
		t.Errorf("Error getting pdf")
	}
	t.Log("Received pdf bytes of length ", len(pdf))
	err = ioutil.WriteFile("testdatapdf1.pdf", pdf, 0644)
	if err != nil {
		t.Errorf("Error writing pdf")
	}
}

func TestGetPNG(t *testing.T) {

	xml, xsl := getXslXml(t)
	png, err := GetPNG(xml, xsl, "testpnguid")
	if err != nil {
		t.Errorf("Error getting png")
	}
	t.Log("Received png bytes of length ", len(png))
	err = ioutil.WriteFile("testdatapng.png", png, 0644)
	if err != nil {
		t.Errorf("Error writing png")
	}
}
