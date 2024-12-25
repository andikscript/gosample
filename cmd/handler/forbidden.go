package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"samplecode/cmd/model"
)

func StatusForbidden(w http.ResponseWriter, maps map[string]interface{}) {
	response := model.Response{
		fmt.Sprintf("%d", http.StatusForbidden),
		http.StatusText(http.StatusForbidden),
		maps,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result, _ := json.Marshal(response)
	w.WriteHeader(http.StatusForbidden)
	w.Write(result)
}
