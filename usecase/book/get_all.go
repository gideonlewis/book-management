package book

import (
	"context"

	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
	"git.teqnological.asia/teq-go/teq-echo/util/myerror"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (u *UseCase) GetAll(ctx context.Context, req *payload.GetAllBookRequest) (*presenter.ListBookResponseWrapper, error) {
	myBooks, err := u.BookRepo.GetAll(ctx, req.Unscoped)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerror.ErrExampleNotFound()
		}

		return nil, myerror.ErrExampleGet(err)
	}

	return &presenter.ListBookResponseWrapper{Books: myBooks}, nil
}
