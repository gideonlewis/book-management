package book

import (
	"context"
	"fmt"
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
			return nil, myerror.ErrBookNotFound()
		}

		return nil, myerror.ErrBookGet(err)
	}

	if req.Title != nil {
		*req.Title = strings.TrimSpace(*req.Title)
		if len(*req.Title) == 0 {
			return nil, myerror.ErrBookInvalidParam("Title")
		}

		myBook.Title = *req.Title
		myBook.Author = *req.Author
		myBook.Price = *req.Price
	}

	myBook.UpdatedBy = teq.Int64(1)
	fmt.Println(myBook)
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
		return nil, myerror.ErrBookUpdate(err)
	}

	return &presenter.BookResponseWrapper{Book: myBook}, nil
}
