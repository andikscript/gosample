package app

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"samplecode/cmd/exception"
	"samplecode/cmd/handler"
	"samplecode/cmd/util"
	"strings"
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	defer exception.CatchUp()

	p := []string{"fileName"}
	trxId := logReceived(r, p)
	initial := true
	body := map[string]interface{}{}

	if r.Method == http.MethodGet {
		fileName := r.FormValue("fileName")

		if fileName == "" {
			body = M{
				"message": "bad request",
			}

			handler.StatusBadRequest(w, body)
			initial = false
		} else {
			fileLocation := filepath.Join("assets", fileName)

			f, err := os.Open(fileLocation)
			if err != nil {
				util.LogError(fmt.Sprintf("Internal server error or '%v'", err))
				splt := strings.Split(fmt.Sprintf("%v", err), ":")

				body = M{
					"message": splt[1][1:],
				}
				handler.StatusInternalServerError(w, body)
				initial = false
			} else {
				contentDisposition := fmt.Sprintf("attachment; filename=%s", filepath.Base(fileLocation))
				w.Header().Set("Content-Disposition", contentDisposition)
				http.ServeFile(w, r, fileLocation)
				initial = false
			}
			defer f.Close()
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
