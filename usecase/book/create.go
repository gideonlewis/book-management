package book

import (
	"context"
	"strings"

	"git.teqnological.asia/teq-go/teq-echo/model"
	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
	"git.teqnological.asia/teq-go/teq-echo/util/myerror"
)

func (u *UseCase) validateCreate(req *payload.CreateBookRequest) error {
	if req.Title == nil {
		return myerror.ErrBookInvalidParam("Title")
	}

	*req.Title = strings.TrimSpace(*req.Title)
	if len(*req.Title) == 0 {
		return myerror.ErrBookInvalidParam("Title")
	}

	return nil
}

func (u *UseCase) Create(
	ctx context.Context,
	req *payload.CreateBookRequest,
) (*presenter.BookResponseWrapper, error) {
	if err := u.validateCreate(req); err != nil {
		return nil, err
	}

	myBook := &model.Book{
		Title:     *req.Title,
		Author:    *req.Author,
		Price:     *req.Price,
		CreatedBy: 1, // must be validate logged Book.
	}
	err := u.BookRepo.Create(ctx, myBook)
	if err != nil {
		return nil, myerror.ErrBookCreate(err)
	}

	return &presenter.BookResponseWrapper{Book: myBook}, nil
}
