package v1

import (
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (r *V1) doLogin(ctx *fiber.Ctx) error {
	u := r.o.AuthURL(ctx.Params("provider"))
	if u == "" {
		return errorResponse(ctx, http.StatusBadRequest, "")
	}

	return ctx.Redirect(u)
}

func (r *V1) doCallback(ctx *fiber.Ctx) error {

	t, err := r.o.GetToken(ctx.Params("provider"), ctx.Query("code"))
	if err != nil {
		log.Printf("token %v", err)
		return errorResponse(ctx, http.StatusBadRequest, "get token fail")
	}

	u, err := r.o.GetUserInfo(ctx.Params("provider"), t)
	if err != nil {
		log.Printf("get user fail %v", err)
		return errorResponse(ctx, http.StatusBadRequest, "get user fail")
	}

	jwt, err := r.o.TokenIssuer(u)
	if err != nil {
		log.Printf("get user fail %v", err)
		return errorResponse(ctx, http.StatusBadRequest, "token issuer fail")
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    jwt,
		Path:     "/",
		SameSite: "none",
		Secure:   true,
		HTTPOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
	})

	return ctx.Redirect("/")
}
