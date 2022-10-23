package main

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"product-service/configs"
	"product-service/db"
	"product-service/handlers"
	"product-service/logger"
	"product-service/repository"
	"product-service/services"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

var config configs.Configuration
var defaultPort string

func init() {
	// load config from config.yaml
	config = configs.LoadConfig()
	defaultPort = strconv.Itoa(config.App.Port)
}

func main() {

	// Setup database -------------------------------------------------------
	db, err := db.NewMySql(config.MySql)
	if err != nil {
		log.Fatal(err)
	}

	// Setup echo -------------------------------------------------------
	e := setupEcho(db)

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
	defer func() {
		db.Close()
		cancel()
	}()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func setupEcho(db *sql.DB) *echo.Echo {

	e := echo.New()
	e.Logger.SetLevel(log.Lvl(config.Log.Level))
	logger.SetLogger(e, config.App.Env) // middleware -- logger

	r := repository.New(db)
	s := services.New(r)
	h := handlers.New(s)

	handlers.SetDefault(e)
	handlers.SetApi(e, h)

	return e
}
