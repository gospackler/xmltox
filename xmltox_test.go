package xmltox

import (
	"io/ioutil"
	"testing"
)

func TestGetPNGFromLink(t *testing.T) {
	converter, err := NewTaskConverter("", "127.0.0.1", []int{2828, 2829})
	if err != nil {
		t.Errorf("Client creation error" + err.Error())
	}
	png, err := converter.GetPNGFromLink("https://google.com")
	if err != nil {
		t.Errorf("Png Convertion" + err.Error())
	}

	ioutil.WriteFile("test.png", png, 0644)
}

/*
func TestGetPDFFromLink(t *testing.T) {
	converter, err := New("")
	if err != nil {
		t.Errorf("Client creation error" + err.Error())
	}
	pdf, err := converter.GetPDFFromLink("http://google.com", 2)
	if err != nil {
		t.Errorf(err.Error())
	}

	ioutil.WriteFile("test.pdf", pdf, 0644)
}
*/
