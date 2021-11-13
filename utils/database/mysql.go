package database

import (
	"context"
	"database/sql"
	"github.com/arcoz0308/arcoz0308.tech/utils"
	_ "github.com/go-sql-driver/mysql"
)

var Database *sql.DB
var ctx = context.TODO()

func Init() {
	db, err := sql.Open("mysql", utils.Config.Database.Dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	Database = db

}
