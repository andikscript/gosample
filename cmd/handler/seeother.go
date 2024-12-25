package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"samplecode/cmd/model"
)

func StatusSeeOther(w http.ResponseWriter, maps map[string]interface{}) {
	response := model.Response{
		fmt.Sprintf("%d", http.StatusSeeOther),
		http.StatusText(http.StatusSeeOther),
		maps,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result, _ := json.Marshal(response)
	w.WriteHeader(http.StatusSeeOther)
	w.Write(result)
}
