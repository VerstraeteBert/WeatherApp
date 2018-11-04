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

func ConnectSQL(host, port, dbname, uname, pass string) (*DB, error) {
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		uname,
		pass,
		host,
		port,
		dbname,
	)

	d, err := sql.Open("mysql", dbSource)
	if err != nil {
		panic(err)
	}

	dbConn.SQL = d

	return dbConn, err
}
