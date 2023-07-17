package book

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/util/myerror"
)

func (u *UseCase) Delete(ctx context.Context, req *payload.DeleteBookRequest) error {
	myBook, err := u.BookRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return myerror.ErrBookNotFound()
		}

		return myerror.ErrBookGet(err)
	}

	err = u.BookRepo.Delete(ctx, myBook, false)
	if err != nil {
		return myerror.ErrBookDelete(err)
	}

	return nil
}
