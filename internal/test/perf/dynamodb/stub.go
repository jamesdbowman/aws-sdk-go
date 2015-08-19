package dynamodb

import (
	"net/http/httptest"
	"net/http"
)

type dbItem struct {
	Key  string
	Data string
}

func successRespServer(resp []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(resp)
	}))
}
