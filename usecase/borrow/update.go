package borrow

import (
	"context"
	"time"

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
			return nil, myerror.ErrBorrowNotFound()
		}

		return nil, myerror.ErrBorrowGet(err)
	}

	if myBorrow.ReturnDate != nil {
		return nil, myerror.ErrBorrowUpdate(err)
	}

	var returnDate time.Time
	if req.ReturnDate != nil {
		returnDate, err = time.Parse(DAY_STANDARD, *req.ReturnDate)
	} else {
		returnDate, err = time.Parse(DAY_STANDARD, time.Now().Format(time.RFC3339))
	}

	if err != nil {
		return nil, myerror.ErrBorrowUpdate(err)
	}

	myBorrow.ReturnDate = &returnDate
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
		return nil, myerror.ErrBorrowUpdate(err)
	}

	return &presenter.BorrowResponseWrapper{Borrow: myBorrow}, nil
}
