package borrow

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
	"git.teqnological.asia/teq-go/teq-echo/util/myerror"
)

func (u *UseCase) GetByID(ctx context.Context, req *payload.GetBorrowByIDRequest) (*presenter.BorrowResponseWrapper, error) {
	fmt.Println("id", req.ID)
	myBorrow, err := u.BorrowRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerror.ErrExampleNotFound()
		}

		return nil, myerror.ErrExampleGet(err)
	}

	return &presenter.BorrowResponseWrapper{Borrow: myBorrow}, nil
}
