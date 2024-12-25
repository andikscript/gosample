package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"samplecode/cmd/model"
)

func StatusAccepted(w http.ResponseWriter, maps map[string]interface{}) {
	response := model.Response{
		fmt.Sprintf("%d", http.StatusAccepted),
		http.StatusText(http.StatusAccepted),
		maps,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result, _ := json.Marshal(response)
	w.WriteHeader(http.StatusAccepted)
	w.Write(result)
}
