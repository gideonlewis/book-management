package user

import (
	"context"

	"git.teqnological.asia/teq-go/teq-echo/codetype"
	"git.teqnological.asia/teq-go/teq-echo/model"
)

type Repository interface {
	Create(ctx context.Context, data *model.User) error
	Update(ctx context.Context, data *model.User) error
	GetByID(ctx context.Context, id int64) (*model.User, error)
	GetAll(ctx context.Context, unscoped bool) ([]model.User, error)
	GetList(
		ctx context.Context,
		search string,
		paginator codetype.Paginator,
		conditions interface{},
		order []string,
	) ([]model.User, int64, error)
	Delete(ctx context.Context, data *model.User, unscoped bool) error
}
