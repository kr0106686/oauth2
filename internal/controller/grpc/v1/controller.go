package v1

import (
	"github.com/go-playground/validator/v10"
	v1 "github.com/kr0106686/oauth2/v2/docs/proto/v1"
	"github.com/kr0106686/oauth2/v2/internal/usecase"
)

// V1 -.
type V1 struct {
	v1.OAuthServer

	t usecase.OAuth
	v *validator.Validate
}
