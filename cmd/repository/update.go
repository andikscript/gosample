package repository

import (
	"fmt"
	"samplecode/cmd/exception"
	"samplecode/cmd/util"
	"time"
)

func UpdateExample(debug bool) {
	defer exception.CatchUp()

	db, err := connectDb(false)
	if err != nil {
		return
	}
	defer db.Close()

	query := "update samplecode.example set example=$1, description=$2, updated=$3 where id = $4"
	args := []interface{}{
		"example updated",
		"description updated",
		time.Now().Format("2006-01-02 15:04:05"),
		"2845252070857459784",
	}

	_, err = db.Exec(query, args...)
	if err != nil {
		util.LogError(fmt.Sprintf("Error when update data on table example or '%v'", err))
		return
	}

	if debug {
		util.LogDebug(fmt.Sprintf("successfully update on table example"))
	}
}
