package model

import (
	"net/http"
	"strings"
	"log"
)

type RequestInfo struct {
	Method  string
	Url     string
	Headers map[string]string
	Body    string
	Id      string
}

func (requestInfo *RequestInfo) GenerateRequest() (*http.Request, error) {

	var request *http.Request
	var err error
	log.Print("GenerateRequest:", requestInfo.Url)
	if len(requestInfo.Body) > 0 {
		request, err = http.NewRequest(requestInfo.Method, requestInfo.Url,
			strings.NewReader(requestInfo.Body))
	} else {
		request, err = http.NewRequest(requestInfo.Method, requestInfo.Url,
			http.NoBody)
	}
	if err != nil {
		return nil, err
	}
	for key, value := range requestInfo.Headers {
		request.Header.Set(key, value)
	}
	return request, nil
}
