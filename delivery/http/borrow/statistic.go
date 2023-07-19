package borrow

import (
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"

	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
)

// GetList Borrows
// @Summary Get Borrow
// @Description Get All Borrow
// @Tags Borrow
// @Accept json
// @Produce json
// @Security AuthToken
// @Success 200 {object} presenter.ListBorrowResponseWrapper
// @Router /borrows/all [get] .
func (r *Route) Statistic(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = payload.StatisticBorrowRequest{}
		resp *presenter.BorrowStatisticResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	resp, err := r.UseCase.Borrow.Statistic(ctx, &req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
