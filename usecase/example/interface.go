package example

import (
	"context"

	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
)

type IUseCase interface {
	Create(ctx context.Context, req *payload.CreateExampleRequest) (*presenter.ExampleResponseWrapper, error)
	Update(ctx context.Context, req *payload.UpdateExampleRequest) (*presenter.ExampleResponseWrapper, error)
	GetByID(ctx context.Context, req *payload.GetByIDRequest) (*presenter.ExampleResponseWrapper, error)
	GetList(ctx context.Context, req *payload.GetListExampleRequest) (*presenter.ListExampleResponseWrapper, error)
	Delete(ctx context.Context, req *payload.DeleteExampleRequest) error
}
