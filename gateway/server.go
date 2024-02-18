package gateway

import (
	"github.com/OLIENTTECH/backend-challenges/internal/cerror"
	"github.com/OLIENTTECH/backend-challenges/ui"
	"github.com/OLIENTTECH/backend-challenges/ui/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func newEchoServer(handler ui.Handler) *echo.Echo {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Validator = validator.NewCustomValidator()
	e.HTTPErrorHandler = cerror.CustomHTTPErrorHandler

	e.GET("/health", handler.HealthCheck().GetHealth)

	v1 := e.Group("/v1")

	// user group
	user := v1.Group("/users")
	{
		userHandler := handler.User()
		user.GET("", userHandler.ListUsers)
		user.POST("", userHandler.PostUser)
		user.GET("/:userID", userHandler.GetUser)
	}

	return e
}
