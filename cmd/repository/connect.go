package repository

import (
	"database/sql"
	"fmt"
	"samplecode/cmd/exception"
	"samplecode/cmd/model"
	"samplecode/cmd/util"
	"samplecode/cmd/util/des"
)

import _ "github.com/lib/pq"

func connectDb(debug bool) (*sql.DB, error) {
	defer exception.CatchUp()

	shared := model.GetConfig()
	password := des.DesDecryption(shared.Database.Password)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		shared.Database.Host, shared.Database.Port, shared.Database.Username, password,
		shared.Database.Dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		util.LogError(fmt.Sprintf("Error when connect db or '%v'", err))
		return nil, err
	}

	if debug {
		util.LogDebug(fmt.Sprintf("successfully connect db"))
	}

	return db, nil
}
