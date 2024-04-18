package cli

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/OLIENTTECH/backend-challenges/internal/cerror"
	"github.com/OLIENTTECH/backend-challenges/ui"
	"github.com/OLIENTTECH/backend-challenges/ui/validator"
)

func newEchoServer(handler ui.Handler) *echo.Echo {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Validator = validator.NewCustomValidator()
	e.HTTPErrorHandler = cerror.CustomHTTPErrorHandler

	e.GET("/health", handler.HealthCheck().GetHealth)
	e.GET("/users", handler.User().ListUsers)

	// v1 := e.Group("/v1")

	// example group
	// example := v1.Group("/examples")
	// {
	// 	exampleHandler := handler.Example()
	// 	example.GET("", exampleHandler.ListUsers)
	// 	example.POST("", exampleHandler.PostUser)
	// 	example.GET("/:exampleID", exampleHandler.GetUser)
	// }

	return e
}
