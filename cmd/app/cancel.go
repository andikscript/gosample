package app

import (
	"fmt"
	"net/http"
	"runtime"
	"samplecode/cmd/exception"
	"samplecode/cmd/handler"
	"samplecode/cmd/util"
	"strings"
	"time"
)

func requestCanceled(w http.ResponseWriter, r *http.Request) {
	defer exception.CatchUp()

	p := []string{"id", "name"}
	trxId := logReceived(r, p)
	initial := true
	body := map[string]interface{}{}

	// check cancel
	runtime.GOMAXPROCS(runtime.NumCPU())
	process := make(chan bool)
	go checkCanceled(process, r, trxId)

	if r.Method == http.MethodGet || r.Method == http.MethodPost {
		if false {
			body = M{
				"message": "bad request",
			}

			handler.StatusBadRequest(w, body)
			initial = false
		}

		body = M{
			"message": "request canceled to samplecode",
		}

		time.Sleep(5 * time.Second)
		process <- true

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

func checkCanceled(ch chan bool, r *http.Request, trxId string) {
	defer exception.CatchUp()

	select {
	case <-r.Context().Done():
		if err := r.Context().Err(); err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "canceled") {
				util.LogInfo(trxId, fmt.Sprintf("- request canceled by user"))
			} else {
				util.LogInfo(trxId, fmt.Sprintf("- unknown error or %s", err.Error()))
			}
		}
	case <-ch:

	}
}
