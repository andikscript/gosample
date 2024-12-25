package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"samplecode/cmd/model"
)

func StatusBadGateway(w http.ResponseWriter, maps map[string]interface{}) {
	response := model.Response{
		fmt.Sprintf("%d", http.StatusBadGateway),
		http.StatusText(http.StatusBadGateway),
		maps,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result, _ := json.Marshal(response)
	w.WriteHeader(http.StatusBadGateway)
	w.Write(result)
}
