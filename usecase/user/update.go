package user

import (
	"context"
	"strings"

	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"git.teqnological.asia/teq-go/teq-echo/model"
	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
	"git.teqnological.asia/teq-go/teq-echo/util/myerror"
)

func (u *UseCase) validateUpdate(ctx context.Context, req *payload.UpdateUserRequest) (*model.User, error) {
	myUser, err := u.UserRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerror.ErrUserNotFound()
		}

		return nil, myerror.ErrUserGet(err)
	}

	if req.Name != nil {
		*req.Name = strings.TrimSpace(*req.Name)
		if len(*req.Name) == 0 {
			return nil, myerror.ErrUserInvalidParam("name")
		}

		myUser.Name = *req.Name
	}

	myUser.UpdatedBy = teq.Int64(1)

	return myUser, nil
}

func (u *UseCase) Update(
	ctx context.Context,
	req *payload.UpdateUserRequest,
) (*presenter.UserResponseWrapper, error) {
	myUser, err := u.validateUpdate(ctx, req)
	if err != nil {
		return nil, err
	}

	err = u.UserRepo.Update(ctx, myUser)
	if err != nil {
		return nil, myerror.ErrUserUpdate(err)
	}

	return &presenter.UserResponseWrapper{User: myUser}, nil
}
