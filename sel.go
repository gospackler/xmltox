package xmltox

// This is the file containing the selenium driver implementation for taking screenshots.
import (
	"errors"
	"path/filepath"

	"sourcegraph.com/sourcegraph/go-selenium"
)

type SelDriver struct {
	xmlFileName string
	xslFileName string
	webDriver   selenium.WebDriver
}

func getAbsPath(fileName string) (string, error) {

	path, err := filepath.Abs(fileName)
	if err != nil {
		return "", err
	}
	return path, nil
}

// Initialize the structure to play with it.
func initSelDriver(xmlFileName string, xslFileName string) (*SelDriver, error) {

	caps := selenium.Capabilities(map[string]interface{}{"browserName": "firefox"})
	driver, err := selenium.NewRemote(caps, "http://localhost:4444/wd/hub")
	if err != nil {
		return nil, errors.New("Failed to open session:" + err.Error())
	}

	xmlAbs, err := getAbsPath(xmlFileName)
	if err != nil {
		return nil, errors.New("Error getting absolute path " + err.Error())
	}

	xslAbs, err := getAbsPath(xslFileName)
	if err != nil {
		return nil, errors.New("Error getting absolute path " + err.Error())
	}

	return &SelDriver{
		xmlFileName: xmlAbs,
		xslFileName: xslAbs,
		webDriver:   driver,
	}, nil
}

func (s *SelDriver) TakeScreenshot() ([]byte, error) {

	err := s.webDriver.Get("file:///" + s.xmlFileName)
	if err != nil {
		return nil, err
	}

	png, err := s.webDriver.Screenshot()
	if err != nil {
		return nil, err
	}

	return png, nil
}

func (s *SelDriver) finishSelDriver() {
	defer s.webDriver.Quit()
}
