package minecraft

import (
	"github.com/gofiber/fiber/v2"
	"sync"
)

func Basic(ctx *fiber.Ctx) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	var pingerr, queryerr error
	go func() {
		defer wg.Done()
	}()
	go func() {
		defer wg.Done()
	}()
}

func Full(ctx *fiber.Ctx) error {

}

func Bedrock(ctx *fiber.Ctx) error {

}
