package user

import (
	"context"

	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
)

type IUseCase interface {
	Create(ctx context.Context, req *payload.CreateUserRequest) (*presenter.UserResponseWrapper, error)
	Update(ctx context.Context, req *payload.UpdateUserRequest) (*presenter.UserResponseWrapper, error)
	GetByID(ctx context.Context, req *payload.GetUserByIDRequest) (*presenter.UserResponseWrapper, error)
	GetList(ctx context.Context, req *payload.GetListUserRequest) (*presenter.ListUserResponseWrapper, error)
	GetAll(ctx context.Context, req *payload.GetAllUserRequest) (*presenter.ListUserResponseWrapper, error)
	Delete(ctx context.Context, req *payload.DeleteUserRequest) error
}
