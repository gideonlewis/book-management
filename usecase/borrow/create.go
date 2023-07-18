package borrow

import (
	"context"
	"fmt"
	"time"

	"git.teqnological.asia/teq-go/teq-echo/model"
	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
	"git.teqnological.asia/teq-go/teq-echo/util/myerror"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const DAY_STANDARD = "2006-01-02T15:04:05Z07:00"

func (u *UseCase) validateCreate(ctx context.Context, req *payload.CreateBorrowRequest) error {
	if req.UserID == nil || req.BookID == nil || req.Quantity == nil {
		return myerror.ErrBorrowInvalidParam("invalid book_id or user_id or quantity")
	}

	if err := u.BorrowRepo.CheckConditions(ctx, req); err != nil {
		return err
	}

	user, err := u.UserRepo.GetByID(ctx, *req.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return myerror.ErrUserNotFound()
		}

		return myerror.ErrUserGet(err)
	}

	if user.ID == 0 {
		return myerror.ErrUserNotFound()
	}

	book, err := u.BookRepo.GetByID(ctx, *req.BookID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return myerror.ErrBookNotFound()
		}

		return myerror.ErrBookGet(err)
	}

	// check quantity
	if book.AvailableQuantity < *req.Quantity {
		err := fmt.Sprintf("quantity not enough, available is %v", book.AvailableQuantity)
		return myerror.ErrBorrowInvalidParam(err)
	}

	return nil
}

func (u *UseCase) Create(
	ctx context.Context,
	req *payload.CreateBorrowRequest,
) (*presenter.BorrowResponseWrapper, error) {
	if err := u.validateCreate(ctx, req); err != nil {
		return nil, err
	}

	borrowDate, err := time.Parse(DAY_STANDARD, time.Now().Format(time.RFC3339))
	if err != nil {
		return nil, myerror.ErrBorrowCreate(err)
	}

	myBorrow := &model.Borrow{
		UserID:     *req.UserID,
		BookID:     *req.BookID,
		Quantity:   *req.Quantity,
		BorrowDate: borrowDate,
		CreatedBy:  1, // must be validate logged Borrow.
	}

	err = u.BorrowRepo.Create(ctx, myBorrow)
	if err != nil {
		return nil, myerror.ErrBorrowCreate(err)
	}

	return &presenter.BorrowResponseWrapper{Borrow: myBorrow}, nil
}
