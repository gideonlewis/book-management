package user

import (
	"context"
	"strings"
	"time"

	"git.teqnological.asia/teq-go/teq-echo/model"
	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
	"git.teqnological.asia/teq-go/teq-echo/util/myerror"
)

func (u *UseCase) validateCreate(req *payload.CreateUserRequest) error {
	if req.Name == nil {
		return myerror.ErrExampleInvalidParam("name")
	}

	*req.Name = strings.TrimSpace(*req.Name)
	if len(*req.Name) == 0 {
		return myerror.ErrExampleInvalidParam("name")
	}

	return nil
}

func (u *UseCase) Create(
	ctx context.Context,
	req *payload.CreateUserRequest,
) (*presenter.UserResponseWrapper, error) {
	if err := u.validateCreate(req); err != nil {
		return nil, err
	}

	myUser := &model.User{
		Name:       *req.Name,
		UserName:   *req.UserName,
		Email:      *req.Email,
		Gender:     *req.Gender,
		Team:       *req.Team,
		DateJoined: time.Now().UTC(),
		CreatedBy:  1, // must be validate logged user.
	}

	err := u.UserRepo.Create(ctx, myUser)
	if err != nil {
		return nil, myerror.ErrExampleCreate(err)
	}

	return &presenter.UserResponseWrapper{User: myUser}, nil
}
