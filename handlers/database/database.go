package database

import (
	"database/sql"
	"github.com/arcoz0308/arcoz0308.tech/handlers/config"
	"github.com/arcoz0308/arcoz0308.tech/handlers/logger"
	"github.com/go-sql-driver/mysql"
	"time"
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

	// for disable unused variable error
	var err error
	DB, err = sql.Open("mysql", cnf.FormatDSN())
	if err != nil {
		logger.AppFatal(true, "mysql", err)
	}
	_, err = Ping()
	if err != nil {
		logger.AppFatal(true, "mysql", err)
	}

}
func Ping() (time.Duration, error) {
	t := time.Now()
	err := DB.Ping()
	if err != nil {
		return -1, err
	}
	return time.Since(t), nil
}
