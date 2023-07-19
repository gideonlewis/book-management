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

	bookStatistics, err := u.BorrowRepo.Statistic(ctx, req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerror.ErrBorrowInvalidParam(err.Error())
		}

		return nil, myerror.ErrBorrowInvalidParam(err.Error())
	}

	//calculate quantum
	var totalBorrowed int64
	for _, statistic := range bookStatistics {
		totalBorrowed += statistic.Quantity
	}

	for _, statistic := range bookStatistics[:10] {
		statistic.Detail = make([]struct {
			Title    string "json:\"title\""
			Quantity int64  "json:\"quantity\""
		}, 1)

		statistic.Detail[0].Title = statistic.Title
		statistic.Detail[0].Quantity = statistic.Quantity
	}

	other := bookStatistics[11]
	other.Detail = make([]struct {
		Title    string "json:\"title\""
		Quantity int64  "json:\"quantity\""
	}, len(bookStatistics)-10)

	other.ID = 0
	other.Title = "Other"
	for i, statistic := range bookStatistics[10:] {
		other.Quantity += statistic.Quantity
		other.NumOfBorrowed += statistic.NumOfBorrowed
		other.Detail[i].Title = statistic.Title
		other.Detail[i].Quantity = statistic.Quantity
	}

	for _, statistic := range bookStatistics {
		statistic.Quantum = float64(statistic.Quantity) / float64(totalBorrowed)
	}

	return &presenter.BorrowStatisticResponseWrapper{Statistics: bookStatistics, Meta: map[string]interface{}{
		"total": len(bookStatistics),
	}}, nil
}
