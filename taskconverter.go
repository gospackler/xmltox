package xmltox

import (
	"fmt"
	"github.com/gospackler/bulldozer"
	"github.com/gospackler/bulldozer/queue"
)

type GenerateType int

const (
	PNG GenerateType = iota
	PDF

	DefaultPDFPages int = 2
)

type OutputData struct {
	Data []byte
	Err  error
}

type InputData struct {
	GenType GenerateType
	Data    []byte
	Link    string
	OutChan chan *OutputData
}

type TaskConverter struct {
	q            *queue.Queue
	workerCount  int
	respChan     chan interface{}
	Input        chan interface{}
	Fin          chan int
	ExitChan     chan int
	DoneChan     chan int
	mapRouteChan chan InputData
}

func NewTaskConverter(workspace, host string, ports []int, workerCount int) (*TaskConverter, error) {
	q := queue.New()
	for _, port := range ports {
		conv, err := New(workspace, host, port)
		if err != nil {
			return nil, err
		}
		fmt.Println("Adding", conv)
		q.Add(conv)
	}
	respChan := make(chan interface{}, workerCount)

	tc := &TaskConverter{
		q:           q,
		workerCount: workerCount,
		respChan:    respChan,
	}
	tc.ExitChan = make(chan int)
	tc.DoneChan = make(chan int)
	freeWorkerChan := bulldozer.InitializeWorkers(workerCount, respChan, tc)
	tc.Input, tc.Fin = bulldozer.Scheduler(freeWorkerChan, tc.ExitChan, respChan, workerCount)
	return tc, nil
}

// This is the embarassingly parallel fuction.
// Accepts InputData and outPuts OutputData
func (t *TaskConverter) Run(inpData interface{}) interface{} {
	convInt := t.q.Remove()

	inp := inpData.(*InputData)
	fmt.Println("Processing input", inp)
	outData := new(OutputData)
	// FIXME Add a wait if instance not available.
	// This should not get hit.
	if convInt == nil {
		fmt.Println("Waiting for free instances ...")
		<-t.DoneChan
		return t.Run(inpData)
	}

	conv := convInt.(*Converter)
	fmt.Println("Converting using", conv)
	var data []byte
	var err error
	if inp.Data != nil {
		switch inp.GenType {
		case PNG:
			data, err = conv.GetPNG(inp.Data)
		case PDF:
			data, err = conv.GetPDF(inp.Data)
		}
	} else {
		switch inp.GenType {
		case PNG:
			data, err = conv.GetPNGFromLink(inp.Link)
		case PDF:
			data, err = conv.GetPDFFromLink(inp.Link, DefaultPDFPages)
		}
	}

	if err != nil {
		outData.Err = err
		inp.OutChan <- outData
		return nil
	}
	outData.Data = data
	t.q.Add(conv)
	select {
	case t.DoneChan <- 1:
	default:
	}
	fmt.Println("Sending data on outChan")
	inp.OutChan <- outData
	return nil
}

func (t *TaskConverter) getOutputFile(inpData *InputData) ([]byte, error) {
	outChan := make(chan *OutputData)
	defer close(outChan)
	inpData.OutChan = outChan
	fmt.Println("Added data", inpData)
	t.Input <- inpData
	output := <-outChan
	if output.Err != nil {
		return nil, output.Err
	}
	return output.Data, nil
}

func (t *TaskConverter) GetPNG(xmlContent []byte) ([]byte, error) {
	inpData := &InputData{
		GenType: PNG,
		Data:    xmlContent,
		Link:    "",
	}
	return t.getOutputFile(inpData)
}

func (t *TaskConverter) GetPNGFromLink(link string) ([]byte, error) {
	inpData := &InputData{
		GenType: PNG,
		Data:    nil,
		Link:    link,
	}
	return t.getOutputFile(inpData)
}

func (t *TaskConverter) GetPDF(xmlContent []byte) ([]byte, error) {
	inpData := &InputData{
		GenType: PDF,
		Data:    xmlContent,
		Link:    "",
	}
	return t.getOutputFile(inpData)
}

func (t *TaskConverter) GetPDFFromLink(link string, noOfPages int) ([]byte, error) {
	inpData := &InputData{
		GenType: PDF,
		Data:    nil,
		Link:    link,
	}
	return t.getOutputFile(inpData)
}

func (t *TaskConverter) Finish() {
	t.Fin <- 1
	<-t.ExitChan
}
