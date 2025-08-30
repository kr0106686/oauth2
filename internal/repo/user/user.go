package user

import (
	"context"

	"github.com/kr0106686/oauth2/v2/internal/entity"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) FirstOrCreate(ctx context.Context, p *entity.User) error {
	return r.db.FirstOrCreate(&entity.User{}, p).Error
}

func (r *Repo) FindByID(ctx context.Context, id uint) (entity.User, error) {
	return gorm.G[entity.User](r.db).Where("id = ?", id).First(ctx)
}

func (r *Repo) Create(ctx context.Context, p *entity.User) error {
	return gorm.G[entity.User](r.db).Create(ctx, p)
}

func (r *Repo) Delete(ctx context.Context, id int) (int, error) {
	return gorm.G[entity.User](r.db).Where("id = ?", id).Delete(ctx)
}

func (r *Repo) Update(ctx context.Context, u *entity.User) (int, error) {
	return gorm.G[*entity.User](r.db).Where("id = ?", u.ID).Updates(ctx, u)
}
