package user

import (
	"context"
	"fmt"

	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
	"git.teqnological.asia/teq-go/teq-echo/util/myerror"
)

func (u *UseCase) GetList(
	ctx context.Context,
	req *payload.GetListUserRequest,
) (*presenter.ListUserResponseWrapper, error) {
	req.Format()

	var (
		order      = make([]string, 0)
		conditions = map[string]interface{}{}
	)

	if req.OrderBy != "" {
		order = append(order, fmt.Sprintf("%s %s", req.OrderBy, req.SortBy))
	}

	if req.CreatedBy != nil && *req.CreatedBy > 0 {
		conditions["created_by"] = req.CreatedBy
	}

	myUsers, total, err := u.UserRepo.GetList(ctx, req.Search, req.Paginator, conditions, order)
	if err != nil {
		return nil, myerror.ErrUserGet(err)
	}

	return &presenter.ListUserResponseWrapper{
		Users: myUsers,
		Meta: map[string]interface{}{
			"page":  req.Paginator.Page,
			"limit": req.Paginator.Limit,
			"total": total,
		},
	}, nil
}
