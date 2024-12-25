package repository

import (
	"fmt"
	"samplecode/cmd/exception"
	"samplecode/cmd/util"
	"time"
)

func InsertExample(debug bool) {
	defer exception.CatchUp()

	db, err := connectDb(false)
	if err != nil {
		return
	}
	defer db.Close()

	query := "insert into samplecode.example values($1,$2,$3,$4,$5)"
	args := []any{time.Now().Format("2006-01-02 15:04:05"), util.RandomTrxId(), "example task",
		"description", time.Now().Format("2006-01-02 15:04:05")}
	_, err = db.Exec(query, args...)
	if err != nil {
		util.LogError(fmt.Sprintf("Error when insert on table example or '%v'", err))
		return
	}

	if debug {
		util.LogDebug(fmt.Sprintf("successfully insert on table example"))
	}
}
