package v1

import (
	"github.com/arcoz0308/arcoz0308.tech/routes/api/v1/arcpaste"
	"github.com/arcoz0308/arcoz0308.tech/routes/api/v1/discord"
	"github.com/gofiber/fiber/v2"
)

func Route(r fiber.Router) {
	r.Get("/discord/users/:id", discord.User)
	arcpaste.Route(r.Group("/arcpaste"))
}
