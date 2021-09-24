package env

import (
	"encoding/json"
	"io/ioutil"
)

var configData map[string]map[string]string

func Load(path string) bool {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return false
	}
	err = json.Unmarshal(file, &configData)
	if err != nil {
		return false
	}
	return true
}

func Get(table, key string) string {
	t, ok := configData[table]
	if !ok {
		return ""
	}
	val, ok := t[key]
	if !ok {
		return ""
	}
	return val
}
