package xmltox

import (
	"io/ioutil"
	"testing"
)

func TestGetPNGFromLink(t *testing.T) {
	converter, err := New("")
	if err != nil {
		t.Errorf("Client creation error" + err.Error())
	}
	png, err := converter.GetPNGFromLink("http://tespeedo.appspot.com")
	if err != nil {
		t.Errorf("Pdf Convertion" + err.Error())
	}

	ioutil.WriteFile("test.png", png, 0644)
}
