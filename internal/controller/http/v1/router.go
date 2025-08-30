package v1

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kr0106686/oauth2/v2/internal/usecase"
)

// NewTranslationRoutes -.
func NewOAuthRoutes(apiV1Group fiber.Router, o usecase.OAuth) {
	r := &V1{o: o, v: validator.New(validator.WithRequiredStructEnabled())}

	{
		apiV1Group.Get("/login/:provider", r.doLogin)
		apiV1Group.Get("/callback/:provider", r.doCallback)
	}
}
