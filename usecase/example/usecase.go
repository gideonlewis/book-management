package example

import (
	"git.teqnological.asia/teq-go/teq-echo/config"
	"git.teqnological.asia/teq-go/teq-echo/repository"
	"git.teqnological.asia/teq-go/teq-echo/repository/example"
	mySES "git.teqnological.asia/teq-go/teq-echo/util/ses"
)

type UseCase struct {
	ExampleRepo example.Repository

	SES mySES.ISES

	Config *config.Config
}

func New(repo *repository.Repository, ses mySES.ISES) IUseCase {
	return &UseCase{
		ExampleRepo: repo.Example,
		SES:         ses,
		Config:      config.GetConfig(),
	}
}
