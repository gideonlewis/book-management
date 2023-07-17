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

const DAY_STANDARD = "2006-01-02T15:04:05Z07:00"

func (u *UseCase) validateCreate(req *payload.CreateUserRequest) error {
	if req.Name == nil {
		return myerror.ErrUserInvalidParam("name")
	}

	*req.Name = strings.TrimSpace(*req.Name)
	if len(*req.Name) == 0 {
		return myerror.ErrUserInvalidParam("name")
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

	joinDate, err := checkJoinDate(req.JoinDate)
	if err != nil {
		return nil, myerror.ErrUserCreate(err)
	}

	myUser := &model.User{
		Name:      *req.Name,
		UserName:  *req.UserName,
		Email:     *req.Email,
		Gender:    *req.Gender,
		Team:      *req.Team,
		JoinDate:  joinDate,
		CreatedBy: 1, // must be validate logged user.
	}

	err = u.UserRepo.Create(ctx, myUser)
	if err != nil {
		return nil, myerror.ErrUserCreate(err)
	}

	return &presenter.UserResponseWrapper{User: myUser}, nil
}

func checkJoinDate(joinDate *string) (time.Time, error) {
	if joinDate == nil {
		return time.Parse(DAY_STANDARD, time.Now().UTC().Format(time.RFC3339))
	}

	return time.Parse(DAY_STANDARD, *joinDate)
}
