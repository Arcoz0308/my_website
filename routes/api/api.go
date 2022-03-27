package api

import (
	v1 "github.com/arcoz0308/arcoz0308.tech/routes/api/v1"
	"github.com/gofiber/fiber/v2"
)

func Route(r fiber.Router) {
	v1.Route(r.Group("/v1"))
}
