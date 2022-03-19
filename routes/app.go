package routes

import (
	"github.com/arcoz0308/arcoz0308.tech/handlers/config"
	v1 "github.com/arcoz0308/arcoz0308.tech/routes/api/v1"
	"github.com/arcoz0308/arcoz0308.tech/utils"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func LoadRoutes() func(ctx *fiber.Ctx) error {
	hosts := make(map[string]*fiber.App, 4)

	// api router
	apiRouter := fiber.New(fiber.Config{AppName: "api"})
	v1.Route(apiRouter.Group("/"))
	hosts["api"] = apiRouter

	// arcpaste router
	arcpasteRouter := fiber.New(fiber.Config{AppName: "arcpaste"})
	hosts["arcpaste"] = arcpasteRouter

	// account router
	accountRouter := fiber.New(fiber.Config{AppName: "account"})
	hosts["account"] = accountRouter

	// main router
	mainRouter := fiber.New(fiber.Config{AppName: "main"})
	hosts["main"] = mainRouter

	return func(ctx *fiber.Ctx) error {
		if ctx.Locals("host_is_ip") != nil || !utils.Prod {
			path := ctx.Path()
			if path == "" || path == "/" {
				return ctx.SendStatus(fiber.StatusNotFound)
			}
			paths := strings.SplitN(path, "/", 3)
			if host, ok := hosts[paths[1]]; ok {
				ctx.Path("/" + paths[2])
				host.Handler()(ctx.Context())
				return nil
			}
			return ctx.SendStatus(fiber.StatusNotFound)
		}
		subdomains := ctx.Subdomains(5)
		if host, ok := hosts[subdomains[0]]; ok {
			host.Handler()(ctx.Context())
			return nil
		}
		if strings.Join(subdomains, ".") == config.Global.Host {
			hosts["mainRouter"].Handler()(ctx.Context())
			return nil
		}
		return ctx.SendStatus(fiber.StatusNotFound)
	}
}
