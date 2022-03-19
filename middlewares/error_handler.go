package middlewares

import (
	error2 "github.com/arcoz0308/arcoz0308.tech/error"
	"github.com/gofiber/fiber/v2"
)

func HandleApiError(ctx fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {

			}
		}
	}()
	err := ctx.Next()
	if err != nil {
		if err, ok := err.(*error2.ApiError); ok {
			return ctx.Status(fiber.StatusInternalServerError).JSON(error2.ApiError{Status: })
		}
	}
}
