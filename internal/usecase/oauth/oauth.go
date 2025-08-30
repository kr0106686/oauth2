package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/kr0106686/oauth2/v2/config"
	v1 "github.com/kr0106686/oauth2/v2/docs/proto/v1"
	"github.com/kr0106686/oauth2/v2/internal/entity"
	"github.com/kr0106686/oauth2/v2/internal/repo"
	"github.com/kr0106686/oauth2/v2/pkg/jwtx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UseCase struct {
	provider repo.Provider
	user     repo.User
	jwt      *jwtx.JWT
}

func New(p repo.Provider, u repo.User, s config.JWT) *UseCase {
	return &UseCase{
		provider: p,
		user:     u,
		jwt:      jwtx.New(s.Secret),
	}
}

func (uc *UseCase) AuthURL(name string) string {
	p := uc.provider.FindProvider(name)
	if p == nil {
		return ""
	}

	u := fmt.Sprintf(
		"%s?client_id=%s&redirect_uri=%s&response_type=code&scope=%s",
		p.Endpoint.AuthURL, p.ClientID,
		p.RedirectURI,
		strings.Join(p.Scopes, " "),
	)

	return u
}

func (uc *UseCase) GetToken(name string, code string) (*entity.Token, error) {
	p := uc.provider.FindProvider(name)
	if p == nil {
		return nil, fmt.Errorf("provider: %s", name)
	}

	v := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"client_id":     {p.ClientID},
		"client_secret": {p.ClientSecret},
		"redirect_uri":  {p.RedirectURI},
	}

	resp, err := http.PostForm(p.Endpoint.TokenURL, v)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("token request failed: %s", string(body))
	}

	var t *entity.Token
	if err := json.NewDecoder(resp.Body).Decode(&t); err != nil {
		return nil, err
	}

	return t, nil
}

func (uc *UseCase) GetUserInfo(name string, t *entity.Token) (*entity.User, error) {
	p := uc.provider.FindProvider(name)
	if p == nil {
		return nil, fmt.Errorf("provider: %s", name)
	}

	req, err := http.NewRequest("GET", p.Endpoint.InfoURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+t.AccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("user info request failed: %s", string(body))
	}

	var u *entity.User
	switch name {
	case "google":
		var g entity.GoogleProfile
		if err := json.NewDecoder(resp.Body).Decode(&g); err != nil {
			return nil, err
		}
		u = &entity.User{
			ProviderID: g.Sub,
			Provider:   "google",
			Email:      g.Email,
			Name:       g.Name,
			Picture:    g.Picture,
		}
	case "kakao":
		var k entity.KakaoProfile
		if err := json.NewDecoder(resp.Body).Decode(&k); err != nil {
			return nil, err
		}
		u = &entity.User{
			ProviderID: fmt.Sprintf("%d", k.ID),
			Provider:   "kakao",
			Email:      k.KakaoAccount.Email,
			Name:       k.KakaoAccount.Profile.Nickname,
			Picture:    k.KakaoAccount.Profile.ProfileImageURL,
		}
	}

	uc.user.FirstOrCreate(context.Background(), u)
	return u, nil
}

func (uc *UseCase) TokenIssuer(u *entity.User) (string, error) {
	return uc.jwt.Issuer(&v1.User{
		Id:         uint64(u.ID),
		ProviderId: u.ProviderID,
		Provider:   u.Provider,
		Name:       u.Name,
		Email:      u.Email,
		Picture:    u.Picture,
		CreatedAt:  timestamppb.New(u.CreatedAt),
		UpdatedAt:  timestamppb.New(u.UpdatedAt),
		DeletedAt:  timestamppb.New(u.DeletedAt.Time),
	})
}

func (uc *UseCase) TokenParser(t string) (*v1.User, error) {
	return uc.jwt.Parser(t)
}
