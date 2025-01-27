package app

import (
	"fmt"
	"net/http"
	"samplecode/cmd/exception"
	"samplecode/cmd/handler"
)

func pathParamHandler(w http.ResponseWriter, r *http.Request) {
	defer exception.CatchUp()

	p := []string{}
	trxId := logReceived(r, p)

	initial := true
	body := map[string]interface{}{}

	if r.Method == http.MethodGet {
		if false {
			body = M{
				"message": "bad request",
			}

			handler.StatusBadRequest(w, body)
			initial = false
		}

		id := r.PathValue("id")
		name := r.PathValue("name")

		body = M{
			"message": "path param from samplecode",
			"body":    fmt.Sprintf("id: %s, name: %s", id, name),
		}

		if initial {
			handler.StatusOk(w, body)
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
