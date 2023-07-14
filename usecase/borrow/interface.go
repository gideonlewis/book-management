package borrow

import (
	"context"

	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
)

type IUseCase interface {
	Create(ctx context.Context, req *payload.CreateBorrowRequest) (*presenter.BorrowResponseWrapper, error)
	Update(ctx context.Context, req *payload.UpdateBorrowRequest) (*presenter.BorrowResponseWrapper, error)
	GetByID(ctx context.Context, req *payload.GetBorrowByIDRequest) (*presenter.BorrowResponseWrapper, error)
	GetList(ctx context.Context, req *payload.GetListBorrowRequest) (*presenter.ListBorrowResponseWrapper, error)
	GetAll(ctx context.Context, req *payload.GetAllBorrowRequest) (*presenter.ListBorrowResponseWrapper, error)
	Delete(ctx context.Context, req *payload.DeleteBorrowRequest) error
}
