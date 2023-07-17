package borrow

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/util/myerror"
)

func (u *UseCase) Delete(ctx context.Context, req *payload.DeleteBorrowRequest) error {
	myBorrow, err := u.BorrowRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return myerror.ErrBorrowNotFound()
		}

		return myerror.ErrBorrowGet(err)
	}

	err = u.BorrowRepo.Delete(ctx, myBorrow, false)
	if err != nil {
		return myerror.ErrBorrowDelete(err)
	}

	return nil
}
