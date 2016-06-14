package xmltox

import (
	"io/ioutil"

	"testing"
)

func TestSel(t *testing.T) {

	selDriver, err := initSelDriver("input/e1_inpatient.xml", "input/CED.XSL")
	if err != nil {
		t.Errorf("Error while initializing selenium driver")
	}
	png, err := selDriver.TakeScreenshot()
	if err != nil {
		t.Errorf("Error getting png")
	}
	t.Log("Received png bytes of length ", len(png))
	err = ioutil.WriteFile("testdata2.png", png, 0644)
	if err != nil {
		t.Errorf("Error writing png")
	}

	defer selDriver.finishSelDriver()
}
