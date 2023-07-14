package book

import (
	"context"

	"git.teqnological.asia/teq-go/teq-echo/codetype"
	"git.teqnological.asia/teq-go/teq-echo/model"
)

type Repository interface {
	Create(ctx context.Context, data *model.Book) error
	Update(ctx context.Context, data *model.Book) error
	GetByID(ctx context.Context, id int64) (*model.Book, error)
	GetAll(ctx context.Context, unscoped bool) ([]model.Book, error)
	GetList(
		ctx context.Context,
		search string,
		paginator codetype.Paginator,
		conditions interface{},
		order []string,
	) ([]model.Book, int64, error)
	Delete(ctx context.Context, data *model.Book, unscoped bool) error
}
