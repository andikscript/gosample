package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"samplecode/cmd/model"
)

func StatusInternalServerError(w http.ResponseWriter, maps map[string]interface{}) {
	response := model.Response{
		fmt.Sprintf("%d", http.StatusInternalServerError),
		http.StatusText(http.StatusInternalServerError),
		maps,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result, _ := json.Marshal(response)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(result)
}
