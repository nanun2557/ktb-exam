package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"product-service/configs"
	"product-service/db"
	"product-service/handlers"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

var defaultPort string

func main() {
	// Setup ----------------------------------------------------------
	e := setupEcho()

	// Start server ----------------------------------------------------------
	go func() {
		if err := e.Start(":" + defaultPort); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func setupEcho() *echo.Echo {

	// load config from config.yaml
	config := configs.LoadConfig()
	defaultPort = strconv.Itoa(config.App.Port)

	e := echo.New()
	e.Logger.SetLevel(log.Lvl(config.Log.Level))

	db, err := db.NewMySql(config.MySql)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// r := repository.New(db)
	// s := services.New(r)
	// h := handlers.New(s)

	handlers.SetDefault(e)

	// e.GET("/", func(c echo.Context) error {
	// 	time.Sleep(5 * time.Second)
	// 	return c.JSON(http.StatusOK, "OK")
	// })
	return e
}
