package authentication

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/arcoz0308/arcoz0308.tech/utils/database"
	"math"
	"time"
)

type cookie struct {
	UserId string
	Value  string
	Type   int
	Expire int
}

var (
	CookieTypeSession = 0
	CookieTypeLong    = 1
)

func GenerateCookie(account *Account, t int) string {
	c := randomBase64String(30)
	var expire int64
	switch t {
	case CookieTypeSession:
		expire = -1
		break
	case CookieTypeLong:
		expire = time.Now().AddDate(0, 0, 30).Unix()
	}
	_, err := database.Database.Exec("INSERT INTO auth_cookies VALUE (?, ?, ?, ?)", account.Id, c, t, expire)
	if err != nil {
		panic(err)
	}
	return c
}
func randomBase64String(l int) string {
	buff := make([]byte, int(math.Ceil(float64(l)/float64(1.33333333333))))
	rand.Read(buff)
	str := base64.RawURLEncoding.EncodeToString(buff)
	return str[:l] // strip the one extra byte we get from half the results.
}
