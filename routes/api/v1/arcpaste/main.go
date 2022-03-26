package arcpaste

import (
	"github.com/gofiber/fiber/v2"
)

func Route(r fiber.Router) {
	r.Get("/:key", paste)
	r.Post("/", newPaste)
	r.Patch("/:key", modifyPaste)
	r.Delete("/:key", deletePaste)
}

func paste(ctx *fiber.Ctx) error {
	return nil
}
func newPaste(ctx *fiber.Ctx) error {
	return nil
}

func modifyPaste(ctx *fiber.Ctx) error {
	return nil
}

func deletePaste(ctx *fiber.Ctx) error {
	return nil
}
