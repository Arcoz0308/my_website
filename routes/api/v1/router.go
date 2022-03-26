package v1

import (
	"github.com/arcoz0308/arcoz0308.tech/routes/api/v1/arcpaste"
	"github.com/gofiber/fiber/v2"
)

func Route(r fiber.Router) {

	arcpaste.Route(r.Group("/arcpaste"))
}
