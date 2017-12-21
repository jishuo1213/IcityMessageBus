package config

import (
	"os"
	"io/ioutil"
	"encoding/json"
)

const REQUEST_QUEUE_NAME string = "request"
const RESPONSE_QUEUE_NAME string = "response"

//const REQUESTER_NUM int = 4

var Config MessageBusConfig

type MessageBusConfig struct {
	MaxProcs int     `json:"go_max_procs"`
	ReadNum  int     `json:"read_count"`
	Hosts    []Hosts `json:"hosts"`
}

type Hosts struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	ToUrl string `json:"toUrl"`
}

func InitConfig() error {
	file, err := os.Open("./config.json")
	if err != nil {
		return err
	}

	configData, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	err = json.Unmarshal(configData, &Config)
	if err != nil {
		return err
	}
	return nil
}
