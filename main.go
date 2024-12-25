package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"samplecode/cmd/app"
	"samplecode/cmd/model"
	"samplecode/cmd/service"
	"samplecode/cmd/util"
	"time"
)

func init() { // initial config
	if model.Shared != nil {
		return
	}

	basePath, err := os.ReadFile(filepath.Join("configs", "config.json"))
	if err != nil {
		util.LogError(fmt.Sprintf("Error when read config or '%v'", err))
	}

	err = json.Unmarshal(basePath, &model.Shared)
	if err != nil {
		util.LogError(fmt.Sprintf("Error when decode config or '%v'", err))
	}
}

func main() {
	router := &service.Middleware{}
	app.RouterHandler(router)

	shared := model.GetConfig()

	address := fmt.Sprintf(":%d", shared.Server.Port)
	server := &http.Server{
		Addr:         address,
		Handler:      router,
		ReadTimeout:  time.Duration(shared.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(shared.Server.WriteTimeout) * time.Second,
	}

	util.LogDebug(fmt.Sprintf("Listening and serving HTTP on %s", address))
	util.LogDebug(fmt.Sprintf("Read Timeout : %s", server.ReadTimeout))
	util.LogDebug(fmt.Sprintf("Write Timeout : %s", server.WriteTimeout))

	err := server.ListenAndServe()
	if err != nil {
		util.LogError(fmt.Sprintf("Error when start server - %v", err.Error()))
	}
}
