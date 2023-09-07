package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func getDbConnStr() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_ADDR"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_TABLENAME"))
}

func Connect() (*bun.DB, error) {
	sqldb, err := sql.Open("mysql", getDbConnStr())
	if err != nil {
		return nil, err
	}

	if err := sqldb.Ping(); err != nil {
		return nil, err
	}

	return bun.NewDB(sqldb, mysqldialect.New()), nil
}
