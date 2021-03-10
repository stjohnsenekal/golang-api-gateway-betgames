package routes

import (
	"api-gateway/config"
	"api-gateway/handlers/betgameshandler"
	"api-gateway/healthutility"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/acme/autocert"
	"net/http"
	"strconv"
	"strings"
)

var configuration config.Configuration

func SetupRoutes(config config.Configuration, healthParam *healthutility.Health) {

	betgameshandler.SetConfiguration(config, healthParam)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")

	b := e.Group("/betgames")
	b.POST("/api", betgameshandler.BetgamesRequestEndoint).Name = "betgames-endpoint"
	b.POST("/usertoken", betgameshandler.GetOrCreateUserTokenEndpoint).Name = "betgames-generate-user-token"

	var port strings.Builder
	port.WriteString(":")
	port.WriteString(strconv.Itoa(config.Port))

	if config.HTTPS {

		e.GET("/", func(c echo.Context) error {
			return c.HTML(http.StatusOK, `
                       <h1>Certificates</h1>
                       <h3>TLS certificates are installed correctly.</h3>
               `)
		})

		e.Logger.Fatal(e.StartTLS(port.String(), config.CERT_file_path, config.KEY_file_path))

	} else {

		e.GET("/", func(c echo.Context) error {
			return c.HTML(http.StatusOK, `
                       <h1>HTTP Active</h1>
                       <h3>Testing Server only.</h3>
               `)
		})

		e.Logger.Fatal(e.Start(port.String()))

	}

}