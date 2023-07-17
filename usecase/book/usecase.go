package book

import (
	"git.teqnological.asia/teq-go/teq-echo/config"
	"git.teqnological.asia/teq-go/teq-echo/repository"
	book "git.teqnological.asia/teq-go/teq-echo/repository/book"
	mySES "git.teqnological.asia/teq-go/teq-echo/util/ses"
)

type UseCase struct {
	BookRepo book.Repository

	SES mySES.ISES

	Config *config.Config
}

func New(repo *repository.Repository, ses mySES.ISES) IUseCase {
	return &UseCase{
		BookRepo: repo.Book,
		SES:      ses,
		Config:   config.GetConfig(),
	}
}
