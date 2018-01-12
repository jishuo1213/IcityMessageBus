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
	"github.com/google/uuid"
)

var server *http.Server
var handler http.Handler

func init() {
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
	toHost, goUrl, err := config.Config.GetToUrl(req.URL.String())

	//log.Println(req.Host, req.URL, req.Method)
	//log.Println(req.URL.Host, req.URL.Path, req.URL.RawPath, req.URL.RawQuery, req.URL.Scheme)

	if err == nil {
		req.Host = toHost
		//req.URL = &url.URL{Path: goUrl}
	} else {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	//req.Host = "http://127.0.0.1:7086/appServer"
	//log.Println(req.Host + req.URL.String())
	body, err := ioutil.ReadAll(req.Body)

	msgDigest := uuid.New().String()
	requestInfo := model.RequestInfo{Method: req.Method, Url: req.Host + goUrl,
		Headers: getRequestHeaders(req), Body: string(body), Id: msgDigest}
	requestBytes, err := utils.EncodeObject(requestInfo)
	if err != nil {
		log.Println(err)
		return
	}

	resChannel := make(chan *model.ResponseData)
	requester.AddResponseChannel(msgDigest, resChannel)
	err = cmsp.PutMessageIntoQueue(config.REQUEST_QUEUE_NAME, requestBytes)
	defer func() {
		requester.RemoveResponseChannel(msgDigest)
		close(resChannel)
	}()
	if err == nil {
		requester.NotifyQueueHasMessage()
		res := <-resChannel
		if res.Header != nil {
			for key, value := range *res.Header {
				if len(value) > 0 {
					w.Header().Set(key, value[0])
					if len(value) > 1 {
						for _, v := range value[1:] {
							w.Header().Add(key, v)
						}
					}
				}
				//if len(value) == 0 {
				//	w.Header().Set(key, value[0])
				//}
			}
		}
		//w.Header().Set("Keep-Alive", strconv.Itoa(0))
		w.Header().Del("Connection")
		w.WriteHeader(res.Code)
		w.Write(res.Body)
	} else {
		w.WriteHeader(500)
		log.Print("put into queue failed")
	}
}

//func NewServerHttp(context iris.Context) {
//	req := context.Request()
//	req.Host = config.Config.Hosts[0].ToUrl
//	log.Println(req.Host + req.URL.String())
//	body, err := ioutil.ReadAll(req.Body)
//
//	msgDigest := uuid.New().String()
//	requestInfo := model.RequestInfo{Method: req.Method, Url: req.Host + req.URL.String(),
//		Headers: getRequestHeaders(req), Body: string(body), Id: msgDigest}
//	requestBytes, err := utils.EncodeObject(requestInfo)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	resChannel := make(chan *model.ResponseData)
//	requester.AddResponseChannel(msgDigest, resChannel)
//	err = cmsp.PutMessageIntoQueue(config.REQUEST_QUEUE_NAME, requestBytes)
//	defer func() {
//		requester.RemoveResponseChannel(msgDigest)
//		close(resChannel)
//	}()
//
//	if err == nil {
//		res := <-resChannel
//		context.StatusCode(res.Code)
//		context.Write(res.Body)
//	} else {
//		context.StatusCode(500)
//		log.Print("put into queue failed")
//	}
//}

func Start(host string, port int) {
	server.Addr = host + ":" + strconv.Itoa(port)

	//app := iris.New()
	//
	//app.Handle("ALL", "/*", irisHandler)
	//
	//app.Build()
	//
	//server.Handler = app

	server.ListenAndServe()
}

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
