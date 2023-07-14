package book

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

func (u *UseCase) validateUpdate(ctx context.Context, req *payload.UpdateBookRequest) (*model.Book, error) {
	myBook, err := u.BookRepo.GetByID(ctx, req.ID)
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

		myBook.Name = *req.Name
	}

	myBook.UpdatedBy = teq.Int64(1)

	return myBook, nil
}

func (u *UseCase) Update(
	ctx context.Context,
	req *payload.UpdateBookRequest,
) (*presenter.BookResponseWrapper, error) {
	myBook, err := u.validateUpdate(ctx, req)
	if err != nil {
		return nil, err
	}

	err = u.BookRepo.Update(ctx, myBook)
	if err != nil {
		return nil, myerror.ErrExampleUpdate(err)
	}

	return &presenter.BookResponseWrapper{Book: myBook}, nil
}
