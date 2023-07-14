package borrow

import (
	"git.teqnological.asia/teq-go/teq-echo/config"
	"git.teqnological.asia/teq-go/teq-echo/repository"
	"git.teqnological.asia/teq-go/teq-echo/repository/borrow"
	mySES "git.teqnological.asia/teq-go/teq-echo/util/ses"
)

type UseCase struct {
	BorrowRepo borrow.Repository

	SES mySES.ISES

	Config *config.Config
}

func New(repo *repository.Repository, ses mySES.ISES) IUseCase {
	return &UseCase{
		BorrowRepo: repo.Borrow,
		SES:        ses,
		Config:     config.GetConfig(),
	}
}
