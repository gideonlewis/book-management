package borrow

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

func (u *UseCase) validateUpdate(ctx context.Context, req *payload.UpdateBorrowRequest) (*model.Borrow, error) {
	myBorrow, err := u.BorrowRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerror.ErrExampleNotFound()
		}

		return nil, myerror.ErrExampleGet(err)
	}

	if req.Name != nil {
		*req.Name = strings.TrimSpace(*req.Name)
		if len(*req.Name) == 0 {
			return nil, myerror.ErrExampleInvalidParam("name")
		}

		myBorrow.Name = *req.Name
	}

	myBorrow.UpdatedBy = teq.Int64(1)

	return myBorrow, nil
}

func (u *UseCase) Update(
	ctx context.Context,
	req *payload.UpdateBorrowRequest,
) (*presenter.BorrowResponseWrapper, error) {
	myBorrow, err := u.validateUpdate(ctx, req)
	if err != nil {
		return nil, err
	}

	err = u.BorrowRepo.Update(ctx, myBorrow)
	if err != nil {
		return nil, myerror.ErrExampleUpdate(err)
	}

	return &presenter.BorrowResponseWrapper{Borrow: myBorrow}, nil
}