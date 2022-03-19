package main

import (
	"crypto/tls"
	"github.com/arcoz0308/arcoz0308.tech/handlers/config"
	"github.com/arcoz0308/arcoz0308.tech/handlers/database"
	"github.com/arcoz0308/arcoz0308.tech/handlers/logger"
	"github.com/arcoz0308/arcoz0308.tech/routes"
	"github.com/arcoz0308/arcoz0308.tech/utils"
	"github.com/gofiber/fiber/v2"
	utils2 "github.com/gofiber/fiber/v2/utils"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"net"
	"os"
)

var Ips map[string]string

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

	//get all local ips
	ips, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, ip := range ips {
		Ips[ip.String()] = ip.String()
	}
	logger.Infof("detected %d ip address", len(Ips))

	// connect to database
	database.Connect()

	// load database prepares of different services
	database.PreparesArcpaste()

	utils.LoadCron()
	utils.StartCron()

	app := fiber.New()

	// redirect to secure
	if utils.Prod {
		app.Use(func(ctx *fiber.Ctx) error {
			hostname := utils2.CopyString(ctx.Hostname())
			// check if already secure
			if ctx.Secure() {
				return ctx.Next()
			}

			_, ok := Ips[hostname]
			if ok {
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
			m := autocert.Manager{
				Prompt:     autocert.AcceptTOS,
				Cache:      autocert.DirCache(config.Cert.Dir),
				HostPolicy: autocert.HostWhitelist(config.Cert.Addrs...),
				Email:      config.Cert.Email,
			}
			cfg := &tls.Config{
				GetCertificate: m.GetCertificate,
				NextProtos: []string{
					"http/1.1", "acme-tls/1",
				},
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
	log.Fatal(app.Listen(port))

}

// GetOutboundIP Get preferred outbound ip of this machine code from
// https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
