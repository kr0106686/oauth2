package http

import (
	"github.com/gofiber/fiber/v2"
	v1 "github.com/kr0106686/oauth2/v2/internal/controller/http/v1"
	"github.com/kr0106686/oauth2/v2/internal/usecase"
)

func NewRouter(app *fiber.App, o usecase.OAuth) {

	// Routers
	apiV1Group := app.Group("/")
	{
		v1.NewOAuthRoutes(apiV1Group, o)
	}
}
