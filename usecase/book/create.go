package book

import (
	"context"
	"strings"

	"git.teqnological.asia/teq-go/teq-echo/client/mysql"
	"git.teqnological.asia/teq-go/teq-echo/model"
	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
	"git.teqnological.asia/teq-go/teq-echo/util/myerror"
	"gorm.io/gorm"
)

func (u *UseCase) validateCreate(req *payload.CreateBookRequest) error {
	if req.Title == nil {
		return myerror.ErrExampleInvalidParam("Title")
	}

	*req.Title = strings.TrimSpace(*req.Title)
	if len(*req.Title) == 0 {
		return myerror.ErrExampleInvalidParam("Title")
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
		Title:             *req.Title,
		Author:            *req.Author,
		Price:             *req.Price,
		TotalQuantity:     1,
		AvailableQuantity: 1,
		CreatedBy:         1, // must be validate logged Book.
	}

	// check book title is existed
	result := mysql.GetDB().Transaction(func(tx *gorm.DB) error {
		var book model.Book
		if err := tx.Where("title = ?", myBook.Title).First(&book).Error; err != nil {
			// if doesnt exist, to create
			myBook.TotalQuantity = 1
			myBook.AvailableQuantity = 1
			err := u.BookRepo.Create(ctx, myBook)
			if err != nil {
				return myerror.ErrExampleCreate(err)
			}

		} else {
			// if existed, to update total_quantity and available_quantity
			if err := tx.Model(&book).UpdateColumns(model.Book{
				TotalQuantity:     book.TotalQuantity + 1,
				AvailableQuantity: book.AvailableQuantity + 1,
			}).Error; err != nil {
				return err
			}
			myBook = &book
		}
		return nil
	})

	if result != nil {
		return &presenter.BookResponseWrapper{Book: &model.Book{}}, nil
	}

	return &presenter.BookResponseWrapper{Book: myBook}, nil
}
