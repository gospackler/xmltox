package xmltox

import (
	"encoding/json"
	"errors"

	"github.com/njasm/marionette_client"
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
	pngData, err := c.GetPNGFromLink(link)
	if err != nil {
		return nil, errors.New("Png conversion error :" + err.Error())
	}
	pdf, err := getPDF(pngData, numOfPages)
	if err != nil {
		return nil, errors.New("Png to pdf conversion error :" + err.Error())
	}
	return pdf, nil
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
	var d = map[string]string{}
	json.Unmarshal([]byte(resp.Value), &d)
	png, err := decodeBase64(d["value"])
	if err != nil {
		errors.New("Decode error" + err.Error())
	}
	return png, nil
}

// Get the signatures right later.
func (c *Converter) GetPNG(xmlContent string) ([]byte, error) {
	return nil, nil
}

func (c *Converter) GetPDF(xmlContent string) ([]byte, error) {
	return nil, nil
}
