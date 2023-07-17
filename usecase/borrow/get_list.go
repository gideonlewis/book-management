package borrow

import (
	"context"
	"fmt"

	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
	"git.teqnological.asia/teq-go/teq-echo/util/myerror"
)

func (u *UseCase) GetList(
	ctx context.Context,
	req *payload.GetListBorrowRequest,
) (*presenter.ListBorrowResponseWrapper, error) {
	req.Format()

	var (
		order      = make([]string, 0)
		conditions = map[string]interface{}{}
	)

	if req.OrderBy != "" {
		order = append(order, fmt.Sprintf("%s %s", req.OrderBy, req.SortBy))
	}

	if req.UserID != nil && *req.UserID > 0 {
		conditions["user_id"] = req.UserID
	}

	if req.CreatedBy != nil && *req.CreatedBy > 0 {
		conditions["created_by"] = req.CreatedBy
	}

	myBorrows, total, err := u.BorrowRepo.GetList(ctx, req.Search, req.Paginator, conditions, order)
	if err != nil {
		return nil, myerror.ErrBorrowGet(err)
	}

	return &presenter.ListBorrowResponseWrapper{
		Borrows: myBorrows,
		Meta: map[string]interface{}{
			"page":  req.Paginator.Page,
			"limit": req.Paginator.Limit,
			"total": total,
		},
	}, nil
}
