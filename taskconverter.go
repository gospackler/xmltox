package xmltox

import (
	"errors"
	"github.com/gospackler/bulldozer/queue"
)

type TaskConverter struct {
	q *queue.Queue
}

func NewTaskConverter(workspace, host string, ports []int) (*TaskConverter, error) {
	queue.New()
	for _, port := range ports {
		conv, err := New(workspace, host, port)
		if err != nil {
			return nil, err
		}
		q.Add(conv)
	}
	return &TaskConverter{
		q: q,
	}, nil
}

type OutputData struct {
	Uuid string
	Data []byte
	Err  error
}

type InputData struct {
	Uuid string
	Data []byte
	url  string
}

// This is the embarassingly parallel fuction.
// Accepts InputData and outPuts OutputData
func (t *TaskConverter) Run(inpData interface{}) interface{} {
	conv := q.Remove()

	// Type assersion as its an interface.
	inpData := inpData.(*InputData)
	outData := new(OutputData)
	if conv != nil {
		outData.Err = errors.New("No firefox instances available in queue")
		return outData
	}

	data, err := conv.GetPNG(inpData.Data)
	if err != nil {
		outData.Err = err
		return outData
	}
	outData.Data = data
	outData.Uuid = inpData.Uuid
	return outData
}
