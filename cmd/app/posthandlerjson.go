package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"samplecode/cmd/exception"
	"samplecode/cmd/handler"
	"samplecode/cmd/model"
)

func postHandlerJson(w http.ResponseWriter, r *http.Request) {
	defer exception.CatchUp()

	p := []string{}
	trxId, jsonBody, isExists := logReceivedPost(r, p)

	initial := true
	body := map[string]interface{}{}

	if r.Method == http.MethodPost {
		req := model.RequestPosting{}
		json.Unmarshal(jsonBody, &req)

		if isExists || req.Id == "" || req.Name == "" || req.Address == "" {
			body = M{
				"message": "bad request",
			}

			handler.StatusBadRequest(w, body)
			initial = false
		}

		request := fmt.Sprintf("%s%s%s", req.Id, req.Name, req.Address)

		if initial {
			body = M{
				"message":  "post from gorestsocket",
				"response": request,
			}

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
