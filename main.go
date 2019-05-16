package main

import (
	"fmt"

	"github.com/devmark88/unireg/handlers"

	"github.com/devmark88/unireg/config"
	"github.com/devmark88/unireg/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()
	c := config.InitSpecs()
	db := models.Init(&c)
	ac := config.AppContext{Config: c, DB: db}
	db.LogMode(true)
	// Middleware
	e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.Logger.SetLevel(0)

	e.GET("/", func(c echo.Context) error {
		return handlers.Hello(c, &ac)
	})
	e.POST("/", func(c echo.Context) error {
		return handlers.AddGraph(c, &ac)
	})

	p := fmt.Sprintf(":%v", c.Port)
	e.Logger.Debugf("connecting to port %v", p)
	e.Logger.Fatal(e.Start(p))
	defer db.Close()
}
