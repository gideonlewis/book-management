package borrow

import (
	"context"
	"time"

	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
	"git.teqnological.asia/teq-go/teq-echo/util/myerror"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const (
	DAY_START = "2015-09-15T14:00:12-00:00"
)

func (u *UseCase) validateStatistic(ctx context.Context, req *payload.StatisticBorrowRequest) error {
	if req.From == nil {
		from := DAY_STANDARD
		req.From = &from
	}

	if req.To == nil {
		to := time.Now().Format(time.RFC3339)
		req.To = &to
	}

	return nil
}

func (u *UseCase) Statistic(ctx context.Context, req *payload.StatisticBorrowRequest) (*presenter.BorrowStatisticResponseWrapper, error) {
	if err := u.validateStatistic(ctx, req); err != nil {
		return nil, myerror.ErrBorrowInvalidParam("invalid from * to date")
	}

	bookStatistic, err := u.BorrowRepo.Statistic(ctx, req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerror.ErrBorrowNotFound()
		}

		return nil, myerror.ErrBorrowGet(err)
	}

	return &presenter.BorrowStatisticResponseWrapper{Books: bookStatistic}, nil
}
