package utility

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	//"github.com/tidwall/gjson"
	"errors"
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

func GetJsonFromFile(fileName string, target interface{}) error {
	if b, err := ioutil.ReadFile(fileName); err == nil {
		return json.Unmarshal(b, &target)
	}
	return errors.New("Error unmarshaling")
}
