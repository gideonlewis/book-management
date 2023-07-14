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
	if req.Name == nil {
		return myerror.ErrExampleInvalidParam("name")
	}

	*req.Name = strings.TrimSpace(*req.Name)
	if len(*req.Name) == 0 {
		return myerror.ErrExampleInvalidParam("name")
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
		Name:      *req.Name,
		Author:    *req.Author,
		Price:     *req.Price,
		CreatedBy: 1, // must be validate logged Book.
	}

	err := u.BookRepo.Create(ctx, myBook)
	if err != nil {
		return nil, myerror.ErrExampleCreate(err)
	}

	return &presenter.BookResponseWrapper{Book: myBook}, nil
}
