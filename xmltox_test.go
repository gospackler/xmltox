package xmltox

import (
	"io/ioutil"
	"testing"
	"time"
)

var converter *TaskConverter

func convert(t *testing.T, link string, fileName string) {
	png, err := converter.GetPNGFromLink(link)
	if err != nil {
		t.Errorf("Png Convertion" + err.Error())
	}
	ioutil.WriteFile(fileName, png, 0644)
}

func set1(t *testing.T) {
	convert(t, "https://google.com", "google.png")
	convert(t, "https://youtube.com", "you.png")
	convert(t, "https://facebook.com", "fb.png")
}

func TestGetPNGFromLink(t *testing.T) {
	var err error
	converter, err = NewTaskConverter("", "127.0.0.1", []int{2828, 2829})
	if err != nil {
		t.Errorf("Client creation error" + err.Error())
	}

	set1(t)
	time.Sleep(time.Second * 10)
	//	converter.Finish()
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
