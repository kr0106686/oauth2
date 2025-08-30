package usecase

import (
	v1 "github.com/kr0106686/oauth2/v2/docs/proto/v1"
	"github.com/kr0106686/oauth2/v2/internal/entity"
)

type (
	OAuth interface {
		AuthURL(name string) string
		GetToken(name string, code string) (*entity.Token, error)
		GetUserInfo(name string, t *entity.Token) (*entity.User, error)

		TokenIssuer(u *entity.User) (string, error)
		TokenParser(t string) (*v1.User, error)
	}
)
