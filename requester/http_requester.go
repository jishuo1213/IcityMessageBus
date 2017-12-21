package requester

import (
	"net/http"
	"io/ioutil"
	"IcityMessageBus/cmsp"
	"IcityMessageBus/config"
	"IcityMessageBus/utils"
	"log"
	"IcityMessageBus/model"
	"sync"
)

var httpClient *http.Client
var responseChannelMap map[string]chan<- *model.ResponseData
var lock sync.RWMutex

func init() {
	httpClient = &http.Client{}
	responseChannelMap = make(map[string]chan<- *model.ResponseData)
	lock = sync.RWMutex{}
}

type ICityRequestBean struct {
	digest  string
	request *http.Request
}

func Start() {
	log.Print("read go count:", config.Config.ReadNum)
	for i := 0; i < config.Config.ReadNum; i++ {
		ch := make(chan *ICityRequestBean, 10)
		go startReadMessageLoop(ch)
		go dealChannel(ch)
	}

}

func dealChannel(requestChannel <-chan *ICityRequestBean) {
	for request := range requestChannel {
		lock.RLock()
		resChannel, found := responseChannelMap[request.digest]
		lock.RUnlock()
		if !found {
			log.Print("not found:", request.digest)
			continue
		}
		code, body, err := sendHttpRequest(request.request)
		log.Print(code)
		if err == nil {
			responseData := model.ResponseData{Body: body, Code: code}
			//err = server.SendResponse(request.digest, code, body)
			if resChannel != nil {
				resChannel <- &responseData
			}
		} else {
			responseData := model.ResponseData{Body: []byte("failed"), Code: 200}
			if resChannel != nil {
				resChannel <- &responseData
			}
			//server.SendResponse(request.digest, 200, []byte("failed"))
		}
		//close(resChannel)
		//RemoveResponseChannel(request.digest)
	}
}

func sendHttpRequest(request *http.Request) (int, []byte, error) {
	//client := &http.Client{}
	resp, err := httpClient.Do(request)
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		log.Print(err)
		return 500, nil, err
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		log.Print(string(body))
		if err != nil {
			log.Print(err)
			return resp.StatusCode, nil, err
		}
		return resp.StatusCode, body, nil
	}
}

func startReadMessageLoop(requestChannel chan<- *ICityRequestBean) {
	for {
		digest, request, err := getRequestFromQueue()
		if err == nil {
			log.Print("err == nil")
			requestChannel <- &ICityRequestBean{digest, request}
		}
	}
}

func getRequestFromQueue() (string, *http.Request, error) {
	msg, err := cmsp.ReadOneMessageFromQueue(config.REQUEST_QUEUE_NAME)

	if err != nil {
		return "", nil, err
	}

	var request model.RequestInfo

	err = utils.DecodeObject(&request, msg)
	log.Print("read one request id:", request.Id)
	if err != nil {
		return "", nil, err
	}

	msgDigest := request.Id

	if err != nil {
		return "", nil, err
	}

	realRequest, err := request.GenerateRequest()

	if err != nil {
		return "", nil, err
	}
	return msgDigest, realRequest, err
}

func AddResponseChannel(digest string, resChannel chan<- *model.ResponseData) {
	lock.Lock()
	responseChannelMap[digest] = resChannel
	lock.Unlock()
}

func RemoveResponseChannel(digest string) () {
	lock.Lock()
	delete(responseChannelMap, digest)
	lock.Unlock()
}
