package book

import (
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"

	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
)

// @Summary Create Book
// @Description Create a Book
// @Tags Book
// @Accept  json
// @Produce json
// @Security AuthToken
// @Param req body payload.CreateBookRequest true "Book info"
// @Success 200 {object} presenter.BookResponseWrapper
// @Router /books [post] .
func (r *Route) Create(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = payload.CreateBookRequest{}
		resp *presenter.BookResponseWrapper
	)
	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	resp, err := r.UseCase.Book.Create(ctx, &req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
