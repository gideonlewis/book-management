package repository

import (
	"context"

	"gorm.io/gorm"

	"git.teqnological.asia/teq-go/teq-echo/repository/book"
	"git.teqnological.asia/teq-go/teq-echo/repository/borrow"
	"git.teqnological.asia/teq-go/teq-echo/repository/example"
	"git.teqnological.asia/teq-go/teq-echo/repository/user"
)

type Repository struct {
	User    user.Repository
	Book    book.Repository
	Borrow  borrow.Repository
	Example example.Repository
}

func New(getClient func(ctx context.Context) *gorm.DB) *Repository {
	return &Repository{
		User:    user.NewPG(getClient),
		Example: example.NewPG(getClient),
	}
}
