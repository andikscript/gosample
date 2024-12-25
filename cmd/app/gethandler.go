package app

import (
	"net/http"
	"samplecode/cmd/exception"
	"samplecode/cmd/handler"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	defer exception.CatchUp()

	p := []string{"id", "name"}
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

		body = M{
			"message": "get from samplecode",
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
