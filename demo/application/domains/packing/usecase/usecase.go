package usecase

import (
	"demo/application/domains/packing/entity"
	"demo/application/domains/packing/models"
	"demo/application/domains/packing/repository"
	"demo/application/model"
	"demo/config"
	"errors"

	"context"
)

type usecase struct {
	lib        *model.Lib
	cfg        *config.Config
	repository repository.IRepository
}

func InitUseCase(
	lib *model.Lib,
	cfg *config.Config,
	repository repository.IRepository,
) IUseCase {
	return &usecase{
		cfg:        cfg,
		lib:        lib,
		repository: repository,
	}
}

func (u *usecase) Read(ctx context.Context, params *models.Params) ([]*entity.User, error) {
	res, err := u.repository.Read(ctx, params)
	if err != nil {
		return nil, errors.New("error reading repository")
	}
	return res, err
}

func (u *usecase) Insert(ctx context.Context, params *entity.User) error {
	if params.Id == 0 {
		return errors.New("id required")
	}
	err := u.repository.Insert(ctx, params)
	if err != nil {
		return errors.New("error inserting repository")
	}
	return err
}

func (u *usecase) Update(ctx context.Context, params *entity.User) error {
	if params.Id == 0 {
		return errors.New("id required")
	}
	err := u.repository.Update(ctx, params)
	if err != nil {
		return errors.New("error inserting repository")
	}
	return err
}
func (u *usecase) Delete(ctx context.Context, params *entity.User) error {
	if params.Id == 0 {
		return errors.New("id required")
	}
	err := u.repository.Delete(ctx, params)
	if err != nil {
		return errors.New("error inserting repository")
	}
	return err
}
