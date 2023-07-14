package grpc

import (
	"git.teqnological.asia/teq-go/teq-echo/config"
	"git.teqnological.asia/teq-go/teq-echo/repository"
	"git.teqnological.asia/teq-go/teq-echo/repository/example"
)

type UseCase struct {
	ExampleRepo example.Repository

	Config *config.Config
}

func New(repo *repository.Repository) IUseCase {
	return &UseCase{
		ExampleRepo: repo.Example,
		Config:      config.GetConfig(),
	}
}
