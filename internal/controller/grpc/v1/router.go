package v1

import (
	"github.com/go-playground/validator/v10"
	v1 "github.com/kr0106686/oauth2/v2/docs/proto/v1"
	"github.com/kr0106686/oauth2/v2/internal/usecase"
	pbgrpc "google.golang.org/grpc"
)

// NewTranslationRoutes -.
func NewOAuthRoutes(app *pbgrpc.Server, t usecase.OAuth) {
	r := &V1{t: t, v: validator.New(validator.WithRequiredStructEnabled())}

	{
		v1.RegisterOAuthServer(app, r)
	}
}
