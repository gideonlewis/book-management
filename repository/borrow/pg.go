package borrow

import (
	"context"
	"errors"
	"fmt"
	"time"

	"git.teqnological.asia/teq-go/teq-echo/codetype"
	"git.teqnological.asia/teq-go/teq-echo/model"
	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
	"gorm.io/gorm"
)

func NewPG(getDB func(ctx context.Context) *gorm.DB) Repository {
	return &pgRepository{getDB}
}

type pgRepository struct {
	getDB func(ctx context.Context) *gorm.DB
}

func (p *pgRepository) Statistic(ctx context.Context, req *payload.StatisticBorrowRequest) ([]*presenter.Statistic, error) {
	var db = p.getDB(ctx)
	if req.Unscoped {
		db.Unscoped()
	}
	// get id & title all book
	var statistics []*presenter.Statistic
	if err := db.Table("books").Select("ID", "Title").Find(&statistics).Error; err != nil {
		return []*presenter.Statistic{}, err
	}

	// calculate num_of_borrowed & total_borrowed
	for _, statistic := range statistics {
		result := db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Table("borrows").Where("book_id = ?", statistic.ID).Count(&statistic.NumOfBorrowed).Error; err != nil {
				return err
			}

			if err := tx.Table("borrows").Select("SUM(quantity) AS quantity").Where("book_id = ?", statistic.ID).Scan(&statistic.Quantity).Error; err != nil {
				return err
			}

			return nil
		})

		if result != nil {
			return []*presenter.Statistic{}, result
		}
	}
	// total borrowed
	var totalBorrowed int64
	if err := db.Table("borrows").Select("SUM(quantity) AS quantity").Scan(&totalBorrowed).Error; err != nil {
		return []*presenter.Statistic{}, err
	}

	// calculate quantum by total_book_borrowed
	for _, statistic := range statistics {
		if statistic.Quantity == nil {
			statistic.Quantum = 0
		} else {
			statistic.Quantum = float64(*statistic.Quantity / totalBorrowed)
		}
	}

	return statistics, nil
}

// func getBookStatistic(db *gorm.DB, unscoped bool) ([]presenter.Statistic, error) {
// 	if unscoped {
// 		db = db.Unscoped()
// 	}

// 	var statistics []presenter.Statistic
// 	if err := db.Table("books").Select("ID", "Title").Find(&statistics).Error; err != nil {
// 		return []presenter.Statistic{}, err
// 	}

// 	for _, statistic := range statistics {
// 		if err := db.Table("borrows").Where("book_id = ?", statistic.ID).Count(&statistic.NumOfBorrowed).Error; err != nil {
// 			return []presenter.Statistic{}, err
// 		}
// 	}

// 	var totalBorrowed int64
// 	if err := db.Table("borrows").Select("quantity").Count(&totalBorrowed).Error; err != nil {
// 		return []presenter.Statistic{}, err
// 	}

// 	return statistics
// }

// func calculateQuantum(db *gorm.DB, bookID string, totalBorrowed int64) float64 {
// 	if err:= db.Table("")
// }

func (p *pgRepository) CheckConditions(ctx context.Context, req *payload.CreateBorrowRequest) error {
	var bookCurr *model.Book
	result := p.getDB(ctx).Transaction(func(tx *gorm.DB) error {
		// check quantity user is borrowed
		var borrowings []model.Borrow
		if err := tx.Where("user_id = ? AND return_date IS NULL", req.UserID).Find(&borrowings).Error; err != nil {
			return err
		}
		if len(borrowings) >= 3 {
			return errors.New("người dùng đã mượn đủ số lượng sách tối đa (3 cuốn)")
		}
		// check available_quantity
		if err := tx.Where("id = ?", req.BookID).First(&bookCurr).Error; err != nil {
			return err
		}
		if bookCurr.AvailableQuantity <= 0 {
			return errors.New("trong kho đã hết sách này")
		}
		// check overdue
		for _, borrowing := range borrowings {
			// Kiểm tra nếu đã quá 2 tuần từ ngày mượn
			if time.Now().Sub(borrowing.BorrowDate).Hours() > 14*24 {
				return errors.New("bạn có sách quá hạn trả, vui lòng trả sách trước khi mượn tiếp")
			}
		}

		return nil
	})

	if result != nil {
		err := fmt.Sprintf("Không thể mượn sách:%s", result)
		return errors.New(err)
	}

	return nil
}

func (p *pgRepository) Create(ctx context.Context, data *model.Borrow) error {
	if err := p.getDB(ctx).Create(data).Error; err != nil {
		return err
	}

	// update available_quantity
	if err := p.getDB(ctx).Table("books").Where("id = ?", data.BookID).UpdateColumn("available_quantity", gorm.Expr("available_quantity - ?", data.Quantity)).Error; err != nil {
		return err
	}

	return nil
}

func (p *pgRepository) Update(ctx context.Context, data *model.Borrow) error {
	return p.getDB(ctx).Save(data).Error
}

func (p *pgRepository) GetByID(ctx context.Context, id int64) (*model.Borrow, error) {
	var borrow model.Borrow

	err := p.getDB(ctx).
		Where("id = ?", id).
		First(&borrow).
		Error

	if err != nil {
		return nil, err
	}

	return &borrow, nil
}

func (p *pgRepository) GetList(
	ctx context.Context,
	search string,
	paginator codetype.Paginator,
	conditions interface{},
	order []string,
) ([]model.Borrow, int64, error) {
	var (
		db     = p.getDB(ctx).Model(&model.Borrow{})
		data   = make([]model.Borrow, 0)
		total  int64
		offset int
	)

	fmt.Println(conditions)
	if conditions != nil {
		db = db.Where(conditions)
	}

	if search != "" {
		db.Where("title LIKE ?", "%"+search+"%")
	}

	for i := range order {
		db = db.Order(order[i])
	}

	if paginator.Page != 1 {
		offset = paginator.Limit * (paginator.Page - 1)
	}

	if paginator.Limit != -1 {
		err := db.Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
	}

	err := db.Limit(paginator.Limit).Offset(offset).Find(&data).Error
	if err != nil {
		return nil, 0, err
	}

	if paginator.Limit == -1 {
		total = int64(len(data))
	}

	return data, total, nil
}

func (p *pgRepository) GetAll(ctx context.Context, unscoped bool) ([]model.Borrow, error) {
	var (
		borrows []model.Borrow
		db      = p.getDB(ctx)
	)

	if unscoped {
		db = db.Unscoped()
	}

	if err := db.Find(&borrows).Error; err != nil {
		return nil, err
	}

	return borrows, nil
}

func (p *pgRepository) Delete(ctx context.Context, data *model.Borrow, unscoped bool) error {
	var db = p.getDB(ctx)

	if unscoped {
		db = db.Unscoped()
	}

	if err := updateAvailableQuantity(db, data); err != nil {
		return err
	}

	return db.Delete(&data).Error
}

func updateAvailableQuantity(db *gorm.DB, data *model.Borrow) error {
	if err := db.Table("books").Where("id = ?", data.BookID).UpdateColumn("available_quantity", gorm.Expr("available_quantity + ?", data.Quantity)).Error; err != nil {
		return err
	}

	return nil
}
