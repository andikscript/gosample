package util

import (
	"fmt"
	"log"
)

//util.LogInfo(util.RandomTrxId(), fmt.Sprintf("send to %s", "http://localhost"))

func LogInfo(trxId, text string) {
	logs := fmt.Sprintf("- [ INFO] - [%s] %s", trxId, text)
	log.Println(logs)
}

func LogDebug(text string) {
	logs := fmt.Sprintf("- [DEBUG] - %s", text)
	log.Println(logs)
}

//util.LogError(fmt.Sprintf("Internal server error or '%v'", err))

func LogError(text string) {
	logs := fmt.Sprintf("- [ERROR] - %s", text)
	log.Println(logs)
}

func LogWarn(text string) {
	logs := fmt.Sprintf("- [ WARN] - %s", text)
	log.Println(logs)
}
