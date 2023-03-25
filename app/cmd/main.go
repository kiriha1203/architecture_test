package main

import (
	"net/http"

	"github.com/kiriha1203/architecture_test/src/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

var e = createMux()

func main() {
	conf, err := config.ConnectMysql()
	if err != nil {
		log.Fatalf("error initializing config: %s", err)
	}
	defer conf.DB.Close()

	log.Info("Successfully connected to MySQL database!")

	e.GET("/", articleIndex)

	e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}

func articleIndex(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
