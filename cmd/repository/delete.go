package repository

import (
	"fmt"
	"samplecode/cmd/exception"
	"samplecode/cmd/util"
)

func DeleteExample(debug bool) {
	defer exception.CatchUp()

	db, err := connectDb(false)
	if err != nil {
		return
	}
	defer db.Close()

	query := "delete from samplecode.example where id = $1"
	args := []interface{}{
		"2845252070857459784",
	}

	_, err = db.Exec(query, args...)
	if err != nil {
		util.LogError(fmt.Sprintf("Error when delete data on table example or '%v'", err))
		return
	}

	if debug {
		util.LogDebug(fmt.Sprintf("successfully delete on table example"))
	}
}
