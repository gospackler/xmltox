package xmltox

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"

	"github.com/njasm/marionette_client"
)

type Converter struct {
	workspace string
	client    *marionette_client.Client
}

// The workspace Directory is the directory where the xmlContent will be dumped to.
func New(workspace string) (*Converter, error) {
	client := marionette_client.NewClient()
	err := client.Connect("127.0.0.1", 2828)
	if err != nil {
		return nil, err
	}

	return &Converter{
		workspace: workspace,
		client:    client,
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

	_, err = c.client.Navigate(link)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Screenshot()
	if err != nil {
		return nil, err
	}
	var d = map[string]string{}
	json.Unmarshal([]byte(resp), &d)
	png, err := decodeBase64(d["value"])
	if err != nil {
		errors.New("Decode error" + err.Error())
	}
	return png, nil
}

func (c *Converter) createXml(xmlContent []byte, fileName string) (string, error) {
	xmlFileName := filepath.Join(c.workspace, fileName)
	err := ioutil.WriteFile(xmlFileName, xmlContent, 0644)
	if err != nil {
		return "", err
	}
	localUrl, err := filepath.Abs(xmlFileName)
	localUrl = "file://" + localUrl
	if err != nil {
		return "", err
	}
	return localUrl, nil
}

// This is the function that handles the temporary file names that come up.
// Need to have a scheduler which interacts to make sure the names are
// unique and can run concurrently.
func (c *Converter) getTempFileName() string {
	return "file.xml"
}

func (c *Converter) GetPNG(xmlContent []byte) ([]byte, error) {
	fileName := c.getTempFileName()
	localUrl, err := c.createXml(xmlContent, fileName)
	if err != nil {
		return nil, err
	}
	return c.GetPNGFromLink(localUrl)
}

func (c *Converter) GetPDF(xmlContent []byte) ([]byte, error) {
	fileName := c.getTempFileName()
	localUrl, err := c.createXml(xmlContent, fileName)
	if err != nil {
		return nil, err
	}
	return c.GetPDFFromLink(localUrl, 2)
}
