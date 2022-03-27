package main

import (
	"crypto/tls"
	"github.com/arcoz0308/arcoz0308.tech/handlers/config"
	"github.com/arcoz0308/arcoz0308.tech/handlers/database"
	"github.com/arcoz0308/arcoz0308.tech/handlers/logger"
	"github.com/arcoz0308/arcoz0308.tech/handlers/logger/console"
	"github.com/arcoz0308/arcoz0308.tech/handlers/redis"
	"github.com/arcoz0308/arcoz0308.tech/routes"
	"github.com/arcoz0308/arcoz0308.tech/utils"
	"github.com/gofiber/fiber/v2"
	utils2 "github.com/gofiber/fiber/v2/utils"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

func init() {
	if os.Getenv("PROD") == "true" {
		utils.Prod = true
	}
}

func main() {
	logger.Infoln("starting website...")
	if utils.Prod {
		logger.Noticeln("working in production mode")
	} else {
		logger.Noticeln("working in development mode")
	}

	// init the config
	config.Init()

	go func() {
		console.LoadConsole()
	}()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		logger.Infoln("connecting to database...")
		t := time.Now()
		// connect to database
		database.Connect()
		logger.Infof("connected to database with success in %s", utils.MsWith2Decimal(time.Since(t)))

		// check ping
		ping, err := database.Ping()
		if err != nil {
			logger.Fatalf(true, "failed to ping database, error : %s", err.Error())
		}
		if ping.Milliseconds() > 100 {
			logger.Warnf("big latency with database detected ( %s ), performance can by are reduce", utils.MsWith2Decimal(ping))
		} else {
			logger.Noticef("pinging database with success, result : %s", utils.MsWith2Decimal(ping))
		}
		// load database prepares of different services
		database.PreparesArcpaste()

		wg.Done()
	}()

	go func() {
		logger.Infoln("connecting to redis...")
		t := time.Now()
		// connect to redis
		redis.Connect()
		logger.Infof("connected to redis with success in %s", utils.MsWith2Decimal(time.Since(t)))

		// check ping
		ping, err := redis.Ping()
		if err != nil {
			logger.Fatalf(true, "failed to ping redis, error : %s", err.Error())
		}
		if ping.Milliseconds() > 30 {
			logger.Warnf("big latency with redis detected ( %s ), performance can by are reduce", utils.MsWith2Decimal(ping))
		} else {
			logger.Noticef("pinging redis with success, result : %s", utils.MsWith2Decimal(ping))
		}

		wg.Done()
	}()
	wg.Wait()
	utils.LoadCron()
	utils.StartCron()

	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	// redirect to secure
	if utils.Prod {
		app.Use(func(ctx *fiber.Ctx) error {
			hostname := utils2.CopyString(ctx.Hostname())
			// check if already secure
			if ctx.Secure() {
				return ctx.Next()
			}
			ok := strings.HasSuffix(strings.ToLower(hostname), config.Global.Host)
			if !ok {
				ctx.Locals("host_is_ip", true)
				return ctx.Next()
			}

			//redirect
			originalURL := utils2.CopyString(ctx.OriginalURL())
			return ctx.Redirect("https://"+hostname+originalURL, fiber.StatusPermanentRedirect)
		})
	}

	app.Use(routes.LoadRoutes())

	// load certs
	if utils.Prod {
		go func() {
			var err error
			cfg := &tls.Config{
				NextProtos: []string{
					"http/1.1", "acme-tls/1",
				},
			}
			cfg.Certificates = make([]tls.Certificate, 1)
			cfg.Certificates[0], err = tls.LoadX509KeyPair(config.Cert.CertFile, config.Cert.Key)
			if err != nil {
				logger.Fatalf(true, "error with ssl certificate, error : %s", err.Error())
			}
			con, err := tls.Listen("tcp", ":443", cfg)
			if err != nil {
				logger.Fatal(true, err)
			}
			log.Fatal(app.Listener(con))
		}()
	}
	var port = ":8080"
	if utils.Prod {
		port = ":80"
	}
	con, err := net.Listen("tcp", port)
	if err != nil {
		logger.Fatal(true, err)
	}
	log.Fatal(app.Listener(con))
}
