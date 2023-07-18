package borrow

import (
	"context"

	"git.teqnological.asia/teq-go/teq-echo/codetype"
	"git.teqnological.asia/teq-go/teq-echo/model"
	"git.teqnological.asia/teq-go/teq-echo/payload"
)

type Repository interface {
	Create(ctx context.Context, data *model.Borrow) error
	Update(ctx context.Context, data *model.Borrow) error
	GetByID(ctx context.Context, id int64) (*model.Borrow, error)
	GetAll(ctx context.Context, unscoped bool) ([]model.Borrow, error)
	GetList(
		ctx context.Context,
		search string,
		paginator codetype.Paginator,
		conditions interface{},
		order []string,
	) ([]model.Borrow, int64, error)
	Delete(ctx context.Context, data *model.Borrow, unscoped bool) error
	CheckConditions(ctx context.Context, data *payload.CreateBorrowRequest) error
}
