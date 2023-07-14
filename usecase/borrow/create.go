package borrow

import (
	"context"
	"strings"
	"time"

	"git.teqnological.asia/teq-go/teq-echo/model"
	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
	"git.teqnological.asia/teq-go/teq-echo/util/myerror"
)

func (u *UseCase) validateCreate(req *payload.CreateBorrowRequest) error {
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
	req *payload.CreateBorrowRequest,
) (*presenter.BorrowResponseWrapper, error) {
	if err := u.validateCreate(req); err != nil {
		return nil, err
	}

	myBorrow := &model.Borrow{
		Name:       *req.Name,
		BorrowName: *req.BorrowName,
		Email:      *req.Email,
		Gender:     *req.Gender,
		Team:       *req.Team,
		DateJoined: time.Now().UTC(),
		CreatedBy:  1, // must be validate logged Borrow.
	}

	err := u.BorrowRepo.Create(ctx, myBorrow)
	if err != nil {
		return nil, myerror.ErrExampleCreate(err)
	}

	return &presenter.BorrowResponseWrapper{Borrow: myBorrow}, nil
}
