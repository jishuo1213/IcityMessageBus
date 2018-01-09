package config

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"github.com/kataras/iris/core/errors"
	"strings"
)

const REQUEST_QUEUE_NAME string = "request"
const RESPONSE_QUEUE_NAME string = "response"

//const REQUESTER_NUM int = 4

var Config MessageBusConfig

type MessageBusConfig struct {
	MaxProcs int     `json:"go_max_procs"`
	ReadNum  int     `json:"read_count"`
	Hosts    []Hosts `json:"hosts"`
	urlTable map[string]Hosts
	//exactUrls   []string
	unExactUrls []string
}

type Hosts struct {
	FromUrl string `json:"fromUrl"`
	ToUrl   string `json:"toHost"`
}

func (config *MessageBusConfig) GetToUrl(fromUrl string) (string, string, error) {

	//host, found := config.urlTable[fromUrl]
	//if found {
	//	return host.ToUrl, "", nil
	//} else {
	if unExactUrl, err := config.getUnExactUrl(fromUrl); err == nil {
		host, found := config.urlTable[unExactUrl]
		if found {
			return host.ToUrl, strings.Replace(fromUrl, unExactUrl, "", 1), nil
		}
	}
	return "", "", errors.New("not found the proxy server" + fromUrl)
	//}
}

func (config *MessageBusConfig) getUnExactUrl(url string) (string, error) {
	for _, unExactUrl := range config.unExactUrls {
		if strings.HasPrefix(url, unExactUrl) {
			return unExactUrl, nil
		}
	}
	return "", errors.New("not found the proxy server" + url)
}

func InitConfig(path string) error {
	file, err := os.Open(path + "/config.json")
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

	if Config.urlTable == nil {
		Config.urlTable = make(map[string]Hosts)
	}

	if Config.unExactUrls == nil {
		Config.unExactUrls = make([]string, 0, 10)
	}

	for _, host := range Config.Hosts {
		Config.unExactUrls = append(Config.unExactUrls, host.FromUrl)
		//if strings.HasSuffix(host.FromUrl, "/") {
		//	host.IsExact = false
		//} else {
		//	host.IsExact = true
		//	//append(Config.exactUrls, host.FromUrl)
		//}
		Config.urlTable[host.FromUrl] = host
	}
	return nil
}
