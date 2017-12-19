package server

import (
	"net/http"
	"time"
	"strconv"
	"IcityMessageBus/cmsp"
	"IcityMessageBus/utils"
	"log"
	"io/ioutil"
	"IcityMessageBus/config"
	"IcityMessageBus/requester"
	"IcityMessageBus/model"
)

var server *http.Server
var handler http.Handler

func init() {
	//responseMap = make(map[string]http.ResponseWriter)
	handler = &RequestParser{}
	server = &http.Server{
		Addr:           ":1213",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

type RequestParser struct {
	Name string ""
}

func (parser *RequestParser) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.Host = config.Config.Hosts[0].ToUrl
	log.Println(req.Host + req.URL.String())
	body, err := ioutil.ReadAll(req.Body)

	requestInfo := model.RequestInfo{Method: req.Method, Url: req.Host + req.URL.String(),
		Headers: getRequestHeaders(req), Body: string(body)}
	requestBytes, err := utils.EncodeObject(requestInfo)
	if err != nil {
		log.Println(err)
		return
	}
	err = cmsp.PutMessageIntoQueue(config.REQUEST_QUEUE_NAME, requestBytes)
	if err == nil {
		msgDigest, err := utils.DigestMessage(requestBytes)
		if err != nil {
			log.Println(err)
			return
		}
		resChannel := make(chan *model.ResponseData)
		requester.AddResponseChannel(msgDigest, resChannel)
		res := <-resChannel
		w.WriteHeader(res.Code)
		w.Write(res.Body)
		//responseMap[msgDigest] = w
	} else {
		//w.Write()
	}

}

func Start(host string, port int) {
	server.Addr = host + ":" + strconv.Itoa(port)
	server.ListenAndServe()
	//http.ListenAndServe(":1213", handler)
}

//func SendResponse(requestDigest string, code int, response []byte) error {
//	defer delete(responseMap, requestDigest)
//	if w, found := responseMap[requestDigest]; found {
//		w.WriteHeader(code)
//		_, err := w.Write(response)
//		if err != nil {
//			return err
//		}
//		return nil
//	} else {
//		return errors.New("")
//	}
//}

func getRequestHeaders(request *http.Request) map[string]string {
	var headers = make(map[string]string)
	for key := range request.Header {
		//log.Println(key, len(request.Header[key]))
		value := request.Header[key]
		if len(value) == 1 {
			headers[key] = value[0]
		}
	}
	return headers
}
