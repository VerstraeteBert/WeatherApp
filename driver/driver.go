package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

func ConnectSQL(databaseUrl string) (*DB, error) {
	dbSource := fmt.Sprintf("%s?charset=utf8",
		databaseUrl,
	)

	d, err := sql.Open("mysql", dbSource)
	if err != nil {
		panic(err)
	}

	dbConn.SQL = d

	return dbConn, err
}
