package repository

import (
	"context"

	"gorm.io/gorm"

	"git.teqnological.asia/teq-go/teq-echo/repository/example"
)

type Repository struct {
	Example example.Repository
}

func New(getClient func(ctx context.Context) *gorm.DB) *Repository {
	return &Repository{
		Example: example.NewPG(getClient),
	}
}
