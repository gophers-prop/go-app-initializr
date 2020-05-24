package config

import (
	"encoding/json"
	"fmt"
	{{ if .Logging.ImportPath }}
	"{{ .Logging.ImportPath }}"
	{{end}}
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var configJSON map[string]interface{}
var configYAML map[string]interface{}

func init() {

	//reading config.json file configuration
	configFileJSON, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Config.ERROR_READING_FILE}} {{end}}
		return
	}
	err = json.Unmarshal([]byte(configFileJSON), &configJSON)
	if err != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Config.ERROR_READING_FILE}} {{end}}
		return
	}
	{{ if .Logging.ImportPath }} {{ .Logging.Messages.Config.SUCCESSFULLY_LOADED}} {{end}}

	//reading config.yaml file configuration
	configFileYAML, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Config.ERROR_READING_FILE}} {{end}}
		return
	}
	err = yaml.Unmarshal([]byte(configFileYAML), &configYAML)
	fmt.Println(configYAML)
	if err != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Config.ERROR_READING_FILE}} {{end}}
		return
	}
	{{ if .Logging.ImportPath }} {{ .Logging.Messages.Config.SUCCESSFULLY_LOADED}} {{end}}
}

//GetConfigJSONByKey  returns value of a key
func GetConfigJSONByKey(key string) interface{} {
	if val, ok := configJSON[key]; ok {
		return val
	}
	return nil
}

//GetConfigJSONByKeyToString returns string value of a key
func GetConfigJSONByKeyToString(key string) string {
	if val, ok := configJSON[key]; ok {
		switch val.(type) {
		case string:
			return fmt.Sprintf("%s", val)
		default:
			return ""
		}
	}
	return ""
}

//GetConfigYAMLByKey  returns value of a key
func GetConfigYAMLByKey(key string) interface{} {
	if val, ok := configYAML[key]; ok {
		return val
	}
	return nil
}

//GetConfigYAMLByKeyToString returns string value of a key
func GetConfigYAMLByKeyToString(key string) string {
	if val, ok := configYAML[key]; ok {
		switch val.(type) {
		case string:
			return fmt.Sprintf("%s", val)
		default:
			return ""
		}
	}
	return ""
}
