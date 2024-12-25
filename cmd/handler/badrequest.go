package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"samplecode/cmd/model"
)

func StatusBadRequest(w http.ResponseWriter, maps map[string]interface{}) {
	response := model.Response{
		fmt.Sprintf("%d", http.StatusBadRequest),
		http.StatusText(http.StatusBadRequest),
		maps,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result, _ := json.Marshal(response)
	w.WriteHeader(http.StatusBadRequest)
	w.Write(result)
}

func BadRequestList(w http.ResponseWriter, maps []map[string]interface{}) {
	response := model.ResponseList{
		fmt.Sprintf("%d", http.StatusBadRequest),
		http.StatusText(http.StatusBadRequest),
		maps,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result, _ := json.Marshal(response)
	w.WriteHeader(http.StatusBadRequest)
	w.Write(result)
}
