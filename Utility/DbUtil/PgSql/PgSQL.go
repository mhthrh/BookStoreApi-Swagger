package PgSql

import (
	"database/sql"
)

func RunQuery(cmd string, db *sql.DB) (*sql.Rows, error) {
	d, err := db.Query(cmd)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func ExecuteCommand(cmd string, db *sql.DB) (*sql.Result, error) {
	r, err := db.Exec(cmd)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
