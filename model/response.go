package model

import "net/http"

type ResponseData struct {
	Body   []byte
	Code   int
	Header *http.Header
}
