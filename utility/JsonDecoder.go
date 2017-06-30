package utility

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
)

func GetJson(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return json.Unmarshal(body, target)
}
