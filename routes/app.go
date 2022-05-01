package routes

import (
	"github.com/arcoz0308/arcoz0308.tech/handlers/config"
	"github.com/arcoz0308/arcoz0308.tech/handlers/logger"
	"github.com/arcoz0308/arcoz0308.tech/routes/api"
	"github.com/arcoz0308/arcoz0308.tech/utils"
	"github.com/gofiber/fiber/v2"
	"strings"
)

var Hosts = make(map[string]*fiber.App, 4)

func Load() {
	// api router
	apiRouter := fiber.New(fiber.Config{AppName: "api"})
	api.Route(apiRouter.Group("/"))
	Hosts["api"] = apiRouter

	// arcpaste router
	arcpasteRouter := fiber.New(fiber.Config{AppName: "arcpaste"})
	arcpasteRouter.Use(func(ctx *fiber.Ctx) error {
		return ctx.SendString("soon...")
	})
	Hosts["arcpaste"] = arcpasteRouter

	// account router
	accountRouter := fiber.New(fiber.Config{AppName: "account"})
	accountRouter.Use(func(ctx *fiber.Ctx) error {
		return ctx.SendString("soon...")
	})
	Hosts["account"] = accountRouter

	// main router
	mainRouter := fiber.New(fiber.Config{AppName: "main"})
	mainRouter.Use(func(ctx *fiber.Ctx) error {
		return ctx.SendString("soon...")
	})
	Hosts["main"] = mainRouter

}

func LoadRoutes() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if ctx == nil {
			logger.Warn("l context detected")
			return nil
		}
		if ctx.Locals("host_is_ip") != nil || !utils.Prod {
			path := ctx.Path()
			if path == "" || path == "/" {
				return ctx.SendStatus(fiber.StatusNotFound)
			}
			paths := strings.SplitN(path, "/", 3)
			if len(paths) <= 1 {
				return ctx.SendStatus(fiber.StatusNotFound)
			}
			if host, ok := Hosts[paths[1]]; ok {
				var p string
				if len(paths) == 2 {
					p = ""
				} else {
					p = paths[2]
				}
				ctx.Path("/" + p)
				host.Handler()(ctx.Context())
				return nil
			}
			return ctx.SendStatus(fiber.StatusNotFound)
		}
		subdomains := ctx.Subdomains(5)
		if host, ok := Hosts[subdomains[0]]; ok {
			host.Handler()(ctx.Context())
			return nil
		}
		if strings.Join(subdomains, ".") == config.Global.Host {
			Hosts["main"].Handler()(ctx.Context())
			return nil
		}
		return ctx.SendStatus(fiber.StatusNotFound)
	}
}
