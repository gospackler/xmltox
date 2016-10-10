package xmltox

import (
	"errors"
	"fmt"
	"github.com/gospackler/bulldozer"
	"github.com/gospackler/bulldozer/queue"
)

type OutputData struct {
	Data []byte
	Err  error
}

type InputData struct {
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
	mapRouteChan chan InputData
}

func NewTaskConverter(workspace, host string, ports []int) (*TaskConverter, error) {
	q := queue.New()
	for _, port := range ports {
		conv, err := New(workspace, host, port)
		if err != nil {
			return nil, err
		}
		fmt.Println("Adding", conv)
		q.Add(conv)
	}
	workerCount := len(ports)
	respChan := make(chan interface{}, workerCount)

	tc := &TaskConverter{
		q:           q,
		workerCount: workerCount,
		respChan:    respChan,
	}
	tc.ExitChan = make(chan int)
	freeWorkerChan := bulldozer.InitializeWorkers(workerCount, respChan, tc)
	tc.Input, tc.Fin = bulldozer.Scheduler(freeWorkerChan, tc.ExitChan, respChan, workerCount)
	return tc, nil
}

// This is the embarassingly parallel fuction.
// Accepts InputData and outPuts OutputData
func (t *TaskConverter) Run(inpData interface{}) interface{} {
	conv := (t.q.Remove()).(*Converter)

	fmt.Println("Trying to run something")
	// Type assersion as its an interface.
	inp := inpData.(*InputData)
	outData := new(OutputData)
	if conv == nil {
		outData.Err = errors.New(" No firefox instances available in queue")
		inp.OutChan <- outData
		return nil
	}

	var data []byte
	var err error
	if inp.Data != nil {
		data, err = conv.GetPNG(inp.Data)
	} else {
		data, err = conv.GetPNGFromLink(inp.Link)
	}

	if err != nil {
		outData.Err = err
		inp.OutChan <- outData
		return nil
	}
	outData.Data = data
	fmt.Println("Sending data on outChan")
	inp.OutChan <- outData
	return nil
}

func (t *TaskConverter) getPNG(inpData *InputData) ([]byte, error) {
	outChan := make(chan *OutputData)
	defer close(outChan)
	inpData.OutChan = outChan
	t.Input <- inpData
	output := <-outChan
	if output.Err != nil {
		return nil, output.Err
	}
	return output.Data, nil
}

func (t *TaskConverter) GetPNG(xmlContent []byte) ([]byte, error) {
	inpData := &InputData{
		Data: xmlContent,
		Link: "",
	}
	return t.getPNG(inpData)
}

func (t *TaskConverter) GetPNGFromLink(link string) ([]byte, error) {
	inpData := &InputData{
		Data: nil,
		Link: link,
	}
	return t.getPNG(inpData)
}

func (t *TaskConverter) Finish() {
	t.Fin <- 1
	<-t.ExitChan
}