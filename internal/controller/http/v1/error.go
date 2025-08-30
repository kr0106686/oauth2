package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kr0106686/oauth2/v2/internal/controller/http/v1/response"
)

func errorResponse(ctx *fiber.Ctx, code int, msg string) error {
	return ctx.Status(code).JSON(response.Error{Error: msg})
}
