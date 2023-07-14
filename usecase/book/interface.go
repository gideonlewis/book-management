package book

import (
	"context"

	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
)

type IUseCase interface {
	Create(ctx context.Context, req *payload.CreateBookRequest) (*presenter.BookResponseWrapper, error)
	Update(ctx context.Context, req *payload.UpdateBookRequest) (*presenter.BookResponseWrapper, error)
	GetByID(ctx context.Context, req *payload.GetBookByIDRequest) (*presenter.BookResponseWrapper, error)
	GetList(ctx context.Context, req *payload.GetListBookRequest) (*presenter.ListBookResponseWrapper, error)
	GetAll(ctx context.Context, req *payload.GetAllBookRequest) (*presenter.ListBookResponseWrapper, error)
	Delete(ctx context.Context, req *payload.DeleteBookRequest) error
}
