package xmltox

import (
	"fmt"
	//	"io/ioutil"
	"testing"
	"time"
)

var converter *TaskConverter

func convert(t *testing.T, link string, fileName string) {
	png, err := converter.GetPDFFromLink(link, 1)
	if err != nil {
		t.Errorf("PDF Convertion" + err.Error())
	}
	t.Log(fileName, " ,", len(png))
	//	ioutil.WriteFile(fileName, png, 0644)
}

func set1(t *testing.T) {
	for i := 0; i < 100; i++ {
		name := fmt.Sprintf("%d.pdf", i)
		go convert(t, "https://google.com", name)
	}
}

func TestGetPNGFromLink(t *testing.T) {
	var err error
	converter, err = NewTaskConverter("", "127.0.0.1", []int{2828, 2829}, 100)
	if err != nil {
		t.Errorf("Client creation error" + err.Error())
	}

	set1(t)
	time.Sleep(time.Second * 60 * 10)
	converter.Finish()
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
