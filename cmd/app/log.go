package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"samplecode/cmd/exception"
	"samplecode/cmd/util"
)

func logReceived(r *http.Request, form []string) string {
	defer exception.CatchUp()

	trxId := util.RandomTrxId()

	formMap := map[string]interface{}{}
	for _, v := range form {
		formMap[v] = r.FormValue(v)
	}
	jsonDataForm, _ := json.Marshal(formMap)

	body := map[string]interface{}{}
	_ = json.NewDecoder(r.Body).Decode(&body)
	jsonData, _ := json.Marshal(body)

	util.LogInfo(trxId, fmt.Sprintf("[received] from -> path :\"%s\", method :\"%s\", header:({\"Content-Type\":\"%v\"}), param :(%s), body :(%s)",
		r.URL.Path, r.Method, r.Header.Get("Content-Type"), string(jsonDataForm), string(jsonData)))
	return trxId
}

func logSend(w http.ResponseWriter, r *http.Request, trxId string, body map[string]interface{}) {
	defer exception.CatchUp()

	jsonData, _ := json.Marshal(body)

	util.LogInfo(trxId, fmt.Sprintf("[  send  ] to -> path :\"%s\", method :\"%s\", header:({\"Content-Type\":\"%v\"}), body :(%s)",
		r.URL.Path, r.Method, w.Header(), string(jsonData)))
}
