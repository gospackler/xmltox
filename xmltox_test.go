package xmltox

import (
	"fmt"
	"sync"
	"testing"
)

var converter *TaskConverter

func convert(t *testing.T, link string, fileName string, wg *sync.WaitGroup) {
	png, err := converter.GetPDFFromLink(link, 1)
	if err != nil {
		t.Errorf("PDF Convertion" + err.Error())
	}
	t.Log(fileName, len(png))
	//	ioutil.WriteFile(fileName, png, 0644)
	wg.Done()
}

func set1(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("%d.pdf", i)
		wg.Add(1)
		go convert(t, "https://google.com", name, &wg)
	}
	wg.Wait()
}

func TestGetPNGFromLink(t *testing.T) {
	var err error
	converter, err = NewTaskConverter("", "127.0.0.1", []int{2828, 2829}, 100)
	if err != nil {
		t.Errorf("Client creation error" + err.Error())
	}

	set1(t)
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
