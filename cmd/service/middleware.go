package service

import (
	"net/http"
	"samplecode/cmd/exception"
)

type M map[string]interface{}

type MList []map[string]interface{}

type Middleware struct {
	http.ServeMux
}

func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer exception.CatchUp()

	m.ServeMux.ServeHTTP(w, r)
}
