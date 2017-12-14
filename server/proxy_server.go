package server

import (
	"net/http"
	"time"
	"strconv"
	"bytes"
	"encoding/binary"
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
	requestBuf := new(bytes.Buffer)

	err := binary.Write(requestBuf, binary.LittleEndian, req)
	if err != nil {
	}
}

func Start(host string, port int) {
	server.Addr = host + ":" + strconv.Itoa(port)
	server.ListenAndServe()
}
