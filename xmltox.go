package xmltox

import (
	"github.com/georgethomas111/marionette_client"
)

type Converter struct {
	xslFileName string
	client      *marionette_client.Client
}

func New(workspaceDirectory string) (*Converter, error) {
	client := marionette_client.NewClient()
	err := client.Connect("127.0.0.1", 2828)
	if err != nil {
		return nil, err
	}
	return &Converter{
		xslFileName: "",
		client:      client,
	}, nil
}

func (c *Converter) GetPDFFromLink(link string, numOfPages int) ([]byte, error) {
	return nil, nil
}

// Returns the base64 encoded version of png
func (c *Converter) GetPNGFromLink(link string) ([]byte, error) {
	_, err := c.client.NewSession("", nil)
	if err != nil {
		return nil, err
	}
	_, err = c.client.Get(link)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Screenshot()
	if err != nil {
		return nil, err
	}
	return []byte(resp), nil
}

// Get the signatures right later.
func (c *Converter) GetPNG(xmlContent string) ([]byte, error) {
	return nil, nil
}

func (c *Converter) GetPDF(xmlContent string) ([]byte, error) {
	return nil, nil
}
