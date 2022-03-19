package database

import (
	"database/sql"
	"github.com/arcoz0308/arcoz0308.tech/handlers/config"
	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	cnf := mysql.Config{
		User:                 config.Database.User,
		Passwd:               config.Database.Passwd,
		Net:                  "tcp",
		Addr:                 config.Database.Addr,
		DBName:               config.Database.DbName,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	DB, err := sql.Open("mysql", cnf.FormatDSN())
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
}
