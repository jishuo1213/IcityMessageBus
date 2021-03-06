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
	"time"
)

type readControl struct {
	isQueueEmpty bool
	readChan     chan int
}

var httpClient *http.Client
var responseChannelMap map[string]chan<- *model.ResponseData
var lock sync.RWMutex
var readChanList []*readControl

func init() {
	httpClient = &http.Client{
		Timeout: time.Second * 20,
	}
	responseChannelMap = make(map[string]chan<- *model.ResponseData)
	lock = sync.RWMutex{}
}

type ICityRequestBean struct {
	digest  string
	request *http.Request
}

func Start() {
	log.Print("read go count:", config.Config.ReadNum)

	readChanList = make([]*readControl, 0, config.Config.ReadNum)
	for i := 0; i < config.Config.ReadNum; i++ {
		ch := make(chan *ICityRequestBean)
		readChan := make(chan int)
		control := readControl{false, readChan}
		readChanList = append(readChanList, &control)
		go startReadMessageLoop(ch, &control)
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
		code, header, body, err := sendHttpRequest(request.request)
		log.Print(code)
		if err == nil {
			responseData := model.ResponseData{Body: body, Header: header, Code: code}
			//err = server.SendResponse(request.digest, code, body)
			if resChannel != nil {
				resChannel <- &responseData
			}
		} else {
			responseData := model.ResponseData{Body: []byte(err.Error()), Code: 500}
			if resChannel != nil {
				resChannel <- &responseData
			}
			//server.SendResponse(request.digest, 200, []byte("failed"))
		}
		//close(resChannel)
		//RemoveResponseChannel(request.digest)
	}
}

func sendHttpRequest(request *http.Request) (int, *http.Header, []byte, error) {
	//client := &http.Client{}
	resp, err := httpClient.Do(request)
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		log.Print(err)
		return 500, nil, nil, err
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		log.Print(string(body))
		if err != nil {
			log.Print(err)
			return resp.StatusCode, nil, nil, err
		}
		return resp.StatusCode, &resp.Header, body, nil
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

func startReadMessageLoop(requestChannel chan<- *ICityRequestBean, control *readControl) {
	for {
		//read := <-readChan
		if control.isQueueEmpty {
			<-control.readChan
		}
		digest, request, err := getRequestFromQueue()
		if err == nil {
			control.isQueueEmpty = false
			log.Print("err == nil")
			requestChannel <- &ICityRequestBean{digest, request}
		} else {
			control.isQueueEmpty = true
		}
	}
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

func NotifyQueueHasMessage() {
	for _, control := range readChanList {
		if control.isQueueEmpty {
			control.readChan <- 1
		}
	}
}
