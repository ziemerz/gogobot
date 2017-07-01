package utility

import (
	"net/http"
	"encoding/xml"
	"io/ioutil"
	"errors"
)

func GetFromXML(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return xml.NewDecoder(resp.Body).Decode(target)
}

func GetXMLFromFile(fileName string, target interface{}) error {
	if b, err := ioutil.ReadFile(fileName); err == nil {
		return xml.Unmarshal(b, &target)
	}
	return errors.New("Error unmarshaling")
}
