package borrow

import (
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"

	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
)

// Create Borrow
// @Summary Create Borrow
// @Description Create a Borrow
// @Tags Borrow
// @Accept  json
// @Produce json
// @Security AuthToken
// @Param req body payload.CreateBorrowRequest true "Borrow info"
// @Success 200 {object} presenter.BorrowResponseWrapper
// @Router /borrows [post] .
func (r *Route) Create(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = payload.CreateBorrowRequest{}
		resp *presenter.BorrowResponseWrapper
	)
	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	resp, err := r.UseCase.Borrow.Create(ctx, &req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
