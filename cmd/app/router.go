package app

import (
	"samplecode/cmd/exception"
	"samplecode/cmd/service"
)

type M map[string]interface{}

type MList []map[string]interface{}

func RouterHandler(router *service.Middleware) {
	defer exception.CatchUp()

	router.HandleFunc("/", mainHandler)

	router.HandleFunc("/get", getHandler)

	router.HandleFunc("/post", postHandler)

	router.HandleFunc("/put", putHandler)

	router.HandleFunc("/patch", patchHandler)

	router.HandleFunc("/delete", deleteHandler)

	router.HandleFunc("/head", headHandler)

	router.HandleFunc("/options", optionsHandler)

	router.HandleFunc("/requestcanceled", requestCanceled)

	// if user access, so all request give cookie
	router.HandleFunc("/cookie", cookieHandler)
}
