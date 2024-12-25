package app

import (
	"net/http"
	"samplecode/cmd/exception"
	"samplecode/cmd/handler"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	defer exception.CatchUp()

	p := []string{}
	trxId := logReceived(r, p)
	initial := true
	body := map[string]interface{}{}

	if r.Method == http.MethodGet || r.Method == http.MethodPost {
		if r.URL.Path == "/" {
			body = M{
				"message": "welcome to samplecode",
			}

			handler.StatusOk(w, body)
			initial = false
		}

		if initial {
			handler.StatusForbidden(w, body)
			initial = false
		}
	}

	if initial {
		handler.StatusMethodNotAllowed(w, M{
			"message": "method not allowed",
		})
	}

	logSend(w, r, trxId, body)
	return
}
