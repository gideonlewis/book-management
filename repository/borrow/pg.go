package borrow

import (
	"context"

	"git.teqnological.asia/teq-go/teq-echo/codetype"
	"git.teqnological.asia/teq-go/teq-echo/model"
	"gorm.io/gorm"
)

func NewPG(getDB func(ctx context.Context) *gorm.DB) Repository {
	return &pgRepository{getDB}
}

type pgRepository struct {
	getDB func(ctx context.Context) *gorm.DB
}

func (p *pgRepository) Create(ctx context.Context, data *model.Borrow) error {
	return p.getDB(ctx).Create(data).Error
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

	if conditions != nil {
		db = db.Where(conditions)
	}

	if search != "" {
		db.Where("name LIKE ?", "%"+search+"%")
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

	return db.Delete(&data).Error
}
