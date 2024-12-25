package repository

import (
	"fmt"
	"samplecode/cmd/exception"
	"samplecode/cmd/model"
	"samplecode/cmd/util"
)

func ReadExample(debug bool) {
	defer exception.CatchUp()

	db, err := connectDb(false)
	if err != nil {
		return
	}
	defer db.Close()

	query := "select * from samplecode.example"
	rows, err := db.Query(query)
	if err != nil {
		util.LogError(fmt.Sprintf("Error when select data on table example or '%v'", err))
		return
	}

	var result []model.Example
	for rows.Next() {
		each := model.Example{}
		err = rows.Scan(&each.Created, &each.Id, &each.Example, &each.Description, &each.Updated)
		if err != nil {
			util.LogError(fmt.Sprintf("Error when append data on table example or '%v'", err))
			return
		}

		result = append(result, each)
	}

	if err := rows.Err(); err != nil {
		util.LogError(fmt.Sprintf("Error rows looping data on table example or '%v'", err))
		return
	}

	if debug {
		util.LogDebug(fmt.Sprintf("successfully select on table example"))
		util.LogDebug(fmt.Sprintf("result: %v", result))
	}
}
