package requester

import (
	"net/http"
	"io/ioutil"
	"IcityMessageBus/cmsp"
	"IcityMessageBus/config"
	"IcityMessageBus/utils"
	"log"
	"time"
	"IcityMessageBus/model"
)

var httpClient *http.Client
var responseChannelMap map[string]chan<- *model.ResponseData

func init() {
	httpClient = &http.Client{}
	responseChannelMap = make(map[string]chan<- *model.ResponseData)
}

type ICityRequestBean struct {
	digest  string
	request *http.Request
}

func Start() {
	for i := 0; i < config.Config.ReadNum; i++ {
		ch := make(chan *ICityRequestBean, 10)
		go startReadMessageLoop(ch)
		go dealChannel(ch)
	}

}

func dealChannel(requestChannel <-chan *ICityRequestBean) {
	for request := range requestChannel {
		code, body, err := sendHttpRequest(request.request)
		log.Print(code)
		resChannel := responseChannelMap[request.digest]
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
		close(resChannel)
		delete(responseChannelMap, request.digest)
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

			requestChannel <- &ICityRequestBean{digest, request}
		} else {
			time.Sleep(time.Millisecond)
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
	if err != nil {
		return "", nil, err
	}

	msgDigest, err := utils.DigestMessage(msg)

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
	responseChannelMap[digest] = resChannel
}
