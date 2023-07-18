package borrow

import (
	"git.teqnological.asia/teq-go/teq-echo/config"
	"git.teqnological.asia/teq-go/teq-echo/repository"
	"git.teqnological.asia/teq-go/teq-echo/repository/book"
	"git.teqnological.asia/teq-go/teq-echo/repository/borrow"
	"git.teqnological.asia/teq-go/teq-echo/repository/user"
	mySES "git.teqnological.asia/teq-go/teq-echo/util/ses"
)

type UseCase struct {
	BorrowRepo borrow.Repository

	UserRepo user.Repository

	BookRepo book.Repository

	SES mySES.ISES

	Config *config.Config
}

func New(repo *repository.Repository, ses mySES.ISES) IUseCase {
	return &UseCase{
		BorrowRepo: repo.Borrow,
		BookRepo:   repo.Book,
		UserRepo:   repo.User,
		SES:        ses,
		Config:     config.GetConfig(),
	}
}
