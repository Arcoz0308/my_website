package authentication

import (
	"crypto/rand"
	argon2 "github.com/arcoz0308/arcoz0308.tech/authentication/crypt/argon2/0x13"
	"strings"
)

var (
	v1 = HashInfo{
		V:          "v1",
		Lib:        "argon2",
		Time:       2,
		Memory:     32 * 1024,
		Threads:    2,
		KeyLen:     32,
		LibVersion: 0x13,
		Before: func(password string) string {
			return password
		},
	}
)

var defaultHashInfo = v1

type HashInfo struct {
	V          string
	Lib        string
	Time       uint32
	Memory     uint32
	Threads    uint8
	KeyLen     uint32
	LibVersion int
	Before     func(password string) string
}

// encodeNewPassword this return the hashed (that are encoded before) and the unique key
func encodePassword(password string, i HashInfo) string {
	if argon2.Version != i.LibVersion {
		panic("invalid argon2 version")
	}
	h := argon2.Key(
		[]byte(password),
		generateRandomSalt(),
		i.Time,
		i.Memory,
		i.Threads,
		i.KeyLen,
	)
	p := i.V + "%%"
	return p + string(h)
}
func generateRandomSalt() []byte {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}
	return salt
}
func checkPassword(password string, hash string) bool {
	v := strings.Split(hash, "%%")[0]
	if v != "v1" {
		panic("invalid version")
	}
	h := encodePassword(password, v1)
	return h == hash
}
