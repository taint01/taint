package repository

import (
	"demo/application/domains/packing/entity"
	"demo/application/domains/packing/models"

	"context"
)

type IRepository interface {
	Read(ctx context.Context, params *models.Params) ([]*entity.User, error)
	Insert(ctx context.Context, params *entity.User) error
	Update(ctx context.Context, params *entity.User) error
	Delete(ctx context.Context, params *entity.User) error
}
