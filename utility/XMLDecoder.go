package utility

import (
	"net/http"
	"encoding/xml"
)

func GetFromXML(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return xml.NewDecoder(resp.Body).Decode(target)
}
