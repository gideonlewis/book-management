package grpc

import (
	"context"

	"git.teqnological.asia/teq-go/teq-echo/model"
)

type IUseCase interface {
	GetByID(ctx context.Context, req *GetByIDRequest) (*model.Example, error)
}
