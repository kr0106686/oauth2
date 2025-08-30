package provider

import (
	"github.com/kr0106686/oauth2/v2/config"
	"github.com/kr0106686/oauth2/v2/internal/entity"
)

var endpoint = map[string]*entity.Endpoint{
	"kakao": {
		AuthURL:  "https://kauth.kakao.com/oauth/authorize",
		TokenURL: "https://kauth.kakao.com/oauth/token",
		InfoURL:  "https://kapi.kakao.com/v2/user/me",
	},
	"google": {
		AuthURL:  "https://accounts.google.com/o/oauth2/auth",
		TokenURL: "https://oauth2.googleapis.com/token",
		InfoURL:  "https://openidconnect.googleapis.com/v1/userinfo",
	},
}

type Repo struct {
	m map[string]*entity.Provider
}

func (r *Repo) FindProvider(name string) *entity.Provider {
	return r.m[name]
}

func New(cfg config.Provider) *Repo {

	m := make(map[string]*entity.Provider)

	m["kakao"] = &entity.Provider{
		ClientID:     cfg.Kakao.ClientID,
		ClientSecret: cfg.Kakao.ClientSecret,
		RedirectURI:  cfg.Kakao.RedirectURI,
		Scopes:       []string{},
		Endpoint:     *endpoint["kakao"],
	}

	m["google"] = &entity.Provider{
		ClientID:     cfg.Google.ClientID,
		ClientSecret: cfg.Google.ClientSecret,
		RedirectURI:  cfg.Google.RedirectURI,
		Scopes:       []string{"profile", "email"},
		Endpoint:     *endpoint["google"],
	}

	return &Repo{m: m}
}
