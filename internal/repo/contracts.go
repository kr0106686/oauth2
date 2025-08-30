package repo

import (
	"context"

	"github.com/kr0106686/oauth2/v2/internal/entity"
)

type (
	Provider interface {
		FindProvider(name string) *entity.Provider
	}

	User interface {
		Create(ctx context.Context, p *entity.User) error
		Delete(ctx context.Context, id int) (int, error)
		FindByID(ctx context.Context, id uint) (entity.User, error)
		FirstOrCreate(ctx context.Context, p *entity.User) error
		Update(ctx context.Context, u *entity.User) (int, error)
	}
)
