package borrow

import (
	"context"
	"errors"
	"fmt"
	"time"

	"git.teqnological.asia/teq-go/teq-echo/client/mysql"
	"git.teqnological.asia/teq-go/teq-echo/model"
	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
	"git.teqnological.asia/teq-go/teq-echo/util/myerror"
	"gorm.io/gorm"
)

const DAY_STANDARD = "2006-01-02T15:04:05Z07:00"

func (u *UseCase) validateCreate(req *payload.CreateBorrowRequest) error {
	if req.UserID == nil || req.BookID == nil {
		return myerror.ErrExampleInvalidParam("invalid book_id or borrower_id")
	}

	return nil
}

func (u *UseCase) Create(
	ctx context.Context,
	req *payload.CreateBorrowRequest,
) (*presenter.BorrowResponseWrapper, error) {
	if err := u.validateCreate(req); err != nil {
		return nil, err
	}

	// check available_quantity by title
	// check overdue
	// check quantity borrowed
	// Mượn sách
	var bookCurrent model.Book
	result := mysql.GetDB().Transaction(func(tx *gorm.DB) error {
		// Kiểm tra số lượng sách đang mượn của người dùng
		userID := req.UserID // ID của người dùng
		var borrowings []model.Borrow
		if err := tx.Where("user_id = ? AND return_date IS NULL", userID).Find(&borrowings).Error; err != nil {
			return err
		}
		if len(borrowings) >= 3 {
			return errors.New("người dùng đã mượn đủ số lượng sách tối đa (3 cuốn)")
			// return fmt.Errorf("người dùng đã mượn đủ số lượng sách tối đa (3 cuốn)")
		}

		// Kiểm tra available_quantity của sách
		bookID := req.BookID // ID của sách
		if err := tx.Where("id = ?", bookID).First(&bookCurrent).Error; err != nil {
			return err
		}
		if bookCurrent.AvailableQuantity <= 0 {
			return errors.New("trong kho đã hết sách này")
			// return fmt.Errorf("sách không còn trong kho")
		}

		// Kiểm tra overdue
		for _, borrowing := range borrowings {
			// Kiểm tra nếu đã quá 2 tuần từ ngày mượn
			if time.Now().Sub(borrowing.BorrowDate).Hours() > 14*24 {
				// return fmt.Errorf("bạn có sách quá hạn trả")
				return errors.New("bạn có sách quá hạn trả, vui lòng trả sách trước khi mượn tiếp")
			}
		}

		return nil
	})

	if result != nil {
		fmt.Println("Bạn không đủ tiêu chí mượn sách: ", result)
		return &presenter.BorrowResponseWrapper{Borrow: &model.Borrow{}}, nil
	}

	//
	borrowDate, err := time.Parse(DAY_STANDARD, time.Now().Format(time.RFC3339))
	if err != nil {
		return nil, myerror.ErrExampleCreate(err)
	}

	myBorrow := &model.Borrow{
		UserID:     *req.UserID,
		BookID:     *req.BookID,
		BorrowDate: borrowDate,
		CreatedBy:  1, // must be validate logged Borrow.
	}

	err = u.BorrowRepo.Create(ctx, myBorrow)
	if err != nil {
		return nil, myerror.ErrExampleCreate(err)
	}

	//Cập nhật available_quantity
	if err := mysql.GetDB().Table("books").Where("id = ?", req.BookID).UpdateColumn("available_quantity", bookCurrent.AvailableQuantity-1).Error; err != nil {
		return &presenter.BorrowResponseWrapper{Borrow: &model.Borrow{}}, err
	}

	return &presenter.BorrowResponseWrapper{Borrow: myBorrow}, nil
}
