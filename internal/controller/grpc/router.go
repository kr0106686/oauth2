package grpc

import (
	v1 "github.com/kr0106686/oauth2/v2/internal/controller/grpc/v1"
	"github.com/kr0106686/oauth2/v2/internal/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// NewRouter -.
func NewRouter(app *grpc.Server, t usecase.OAuth) {
	{
		v1.NewOAuthRoutes(app, t)
	}

	reflection.Register(app)
}
