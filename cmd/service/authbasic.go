package service

import (
	"fmt"
	"net/http"
	"samplecode/cmd/handler"
	"samplecode/cmd/model"
	"samplecode/cmd/util"
)

func AuthBasic(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		basicAuth := model.GetBasicAuth()

		username, password, ok := r.BasicAuth()
		body := M{}

		if !ok {
			body = M{
				"message": "something went wrong",
			}

			handler.StatusUnauthorized(w, body)
			util.LogError(fmt.Sprintf("something went wrong for auth"))
			return
		}

		if isValid := (username == basicAuth.BasicAuth.Username) &&
			(password == basicAuth.BasicAuth.Password); !isValid {
			body = M{
				"message": "username or password is wrong",
			}

			handler.StatusUnauthorized(w, body)
			util.LogError(fmt.Sprintf("username or password is wrong"))
			return
		}

		nextHandler.ServeHTTP(w, r)
	})
}
