package logger

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func SetLogger(e *echo.Echo, env string) {
	var logger *zap.Logger
	var err error

	if env == "development" {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		panic(err)
	}

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogMethod:    true,
		LogRequestID: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			requestUUID := v.RequestID
			if requestUUID == "" {
				requestUUID = uuid.NewString()
			}

			logger.Info("request",
				zap.String("URI", v.URI),
				zap.String("method", v.Method),
				zap.Int("status", v.Status),
				zap.String("request_id", requestUUID),
			)

			return nil
		},
	}))

}
