package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"samplecode/cmd/model"
)

func StatusOk(w http.ResponseWriter, maps map[string]interface{}) {
	response := model.Response{
		fmt.Sprintf("%d", http.StatusOK),
		http.StatusText(http.StatusOK),
		maps,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func StatusOkList(w http.ResponseWriter, maps []map[string]interface{}) {
	response := model.ResponseList{
		fmt.Sprintf("%d", http.StatusOK),
		http.StatusText(http.StatusOK),
		maps,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
