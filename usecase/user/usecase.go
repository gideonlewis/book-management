package user

import (
	"git.teqnological.asia/teq-go/teq-echo/config"
	"git.teqnological.asia/teq-go/teq-echo/repository"
	user "git.teqnological.asia/teq-go/teq-echo/repository/user"
	mySES "git.teqnological.asia/teq-go/teq-echo/util/ses"
)

type UseCase struct {
	UserRepo user.Repository

	SES mySES.ISES

	Config *config.Config
}

func New(repo *repository.Repository, ses mySES.ISES) IUseCase {
	return &UseCase{
		UserRepo: repo.User,
		SES:      ses,
		Config:   config.GetConfig(),
	}
}
