package minecraft

import (
	"github.com/gofiber/fiber/v2"
	"net"
	"time"
)

const (
	DefaultPortJava    = 25565
	DefaultPortBedrock = 19132
)

func Basic(ctx *fiber.Ctx) error {

}

func Full(ctx *fiber.Ctx) error {

}

func Bedrock(ctx *fiber.Ctx) error {

}

type basic struct {
	MOTD     string `json:"motd"`
	GameType string `json:"game_type"`
	Map      string `json:"map"`
	Players  struct {
		Online int
		Max    int
	}
	Host struct {
		Ip   string
		Port int
	}
}
type full struct {
}
type bedrock struct {
}

func QueryBasic(adr string) (*basic, error) {
	con, err := net.Dial("udp", adr)
	if err != nil {
		return nil, err
	}
	defer con.Close()
	err = con.SetDeadline(time.Now().Add(time.Second * 5))
	if err != nil {
		return nil, err
	}
}

func QueryFull(adr string) (interface{}, error) {
	con, err := net.Dial("udp", adr)
	if err != nil {
		return nil, err
	}
	defer con.Close()
	err = con.SetDeadline(time.Now().Add(time.Second * 5))
	if err != nil {
		return nil, err
	}
}

func QueryBedrock(adr string) (interface{}, error) {
	con, err := net.Dial("udp", adr)
	if err != nil {
		return nil, err
	}
	defer con.Close()
	err = con.SetDeadline(time.Now().Add(time.Second * 5))
	if err != nil {
		return nil, err
	}
}
