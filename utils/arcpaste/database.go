package arcpaste

import (
	"github.com/arcoz0308/arcoz0308.tech/utils/database"
)

type PasteStruct struct {
	Key      string
	UserId   string
	Raw      string
	Language string
	Expire   int
	Password string
}

func Data(key string) (PasteStruct, error) {
	var r = PasteStruct{}
	row := database.Database.QueryRow("SELECT * FROM paste WHERE id =?", key)
	err := row.Scan(&r.Key, &r.UserId, &r.Raw, &r.Language, &r.Expire, &r.Password)
	return r, err
}
