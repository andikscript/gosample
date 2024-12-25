package app

import (
	"net/http"
	"samplecode/cmd/exception"
	"samplecode/cmd/handler"
	"samplecode/cmd/util"
	"strings"
	"time"
)

func cookieHandler(w http.ResponseWriter, r *http.Request) {
	defer exception.CatchUp()

	p := []string{"cookie"}
	trxId := logReceived(r, p)

	initial := true
	body := map[string]interface{}{}

	if r.Method == http.MethodGet {
		cookie := r.FormValue("cookie")
		c := &http.Cookie{}
		cookieData := "CookieData"

		if cookie == "true" {
			if storedCookie, err := r.Cookie(cookieData); err == nil {
				c = storedCookie
			}

			if c.Value == "" {
				c = &http.Cookie{}
				c.Name = cookieData
				c.Value = util.RandomNumberString(64)
				c.Expires = time.Now().Add(10 * time.Minute)
				http.SetCookie(w, c)
			}
		}

		if cookie == "false" {
			c = &http.Cookie{}
			c.Name = cookieData
			c.Expires = time.Unix(0, 0)
			c.MaxAge = -1
			http.SetCookie(w, c)
		}

		// if empty or ""
		if strings.TrimSpace(cookie) == "" || cookie == "" {
			body = M{
				"message": "bad request",
			}

			handler.StatusBadRequest(w, body)
			initial = false
		}

		body = M{
			"message": "cookie handler from samplecode",
			"cookie":  c.Value,
		}

		if initial {
			handler.StatusOk(w, body)
			initial = false
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
