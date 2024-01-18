package init

import (
	"demo/application/domains/packing/delivery/api/handler"
	"demo/application/domains/packing/repository"
	"demo/application/domains/packing/usecase"
	"demo/application/model"
	"demo/config"
)

type Init struct {
	Repository repository.IRepository
	Usecase    usecase.IUseCase
	Handler    handler.Handler
}

func NewInit(
	lib *model.Lib,
	cfg *config.Config,
) *Init {
	repository := repository.InitRepo(lib)
	usecase := usecase.InitUseCase(lib, cfg, repository)
	handler := handler.InitHandler(lib, cfg, usecase)
	return &Init{
		Repository: repository,
		Usecase:    usecase,
		Handler:    handler,
	}
}
