package authentication

import (
	"database/sql"
	"errors"
	"github.com/arcoz0308/arcoz0308.tech/utils/database"
)

type Account struct {
	Id       string
	Username string
	Mail     string
	Avatar   string
}

var (
	ErrUserDontExist      = errors.New("user don't exist")
	ErrPasswordAreInvalid = errors.New("the password is invalid")
	ErrVerifyEmail        = errors.New("the email don't are verified")
)

func GetAccount(name string, password string, checkPasswordAllTime bool) (*Account, error) {
	row := database.Database.QueryRow("SELECT * FROM users WHERE (username=?) OR (email=?)", name, name)
	a := &Account{}
	var verified bool
	var h string
	err := row.Scan(&a.Id, &a.Username, &a.Mail, &h, &verified, &a.Avatar)
	if err != nil {
		if err == sql.ErrNoRows {
			if checkPasswordAllTime {
				encodePassword(password, defaultHashInfo)
				return nil, ErrUserDontExist
			}
			return nil, ErrUserDontExist
		}
		panic(err)
		return nil, err
	}
	if !checkPassword(password, h) {
		return nil, ErrPasswordAreInvalid
	}
	if err != nil {
		panic(err)
		return a, err
	}
	return a, nil
}
