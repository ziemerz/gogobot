package utility

import (
	"io/ioutil"
	"encoding/json"
	"reflect"
	"strings"
)

type Keys struct {
	Cat string `json:"cat"`
	Dog string `json:"giphy"`
}

func loadKeys() Keys{
	var keys Keys
	if b, err := ioutil.ReadFile("apikeys.json"); err == nil {
		json.Unmarshal(b, &keys)
	}
	return keys
}

func GetKey(key string) string{
	keys := loadKeys()
	st := reflect.ValueOf(keys).FieldByName(strings.Title(key))
	return st.String()
}
