package service

import (
	"net/http"
	"samplecode/cmd/exception"
)

type M map[string]interface{}

type MList []map[string]interface{}

type Middleware struct {
	http.ServeMux
	middleware []func(next http.Handler) http.Handler
}

func (m *Middleware) RegisterMiddleware(next func(http.Handler) http.Handler) {
	m.middleware = append(m.middleware, next)
}

func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer exception.CatchUp()

	var current http.Handler = &m.ServeMux
	for _, h := range m.middleware {
		current = h(current)
	}

	current.ServeHTTP(w, r)
}
