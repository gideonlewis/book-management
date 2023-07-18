package book

import (
	"fmt"

	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"

	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
)

// @Summary Get Book
// @Description Get all Book by unscoped
// @Tags Book
// @Accept json
// @Produce json
// @Security AuthToken
// @Success 200 {object} presenter.ListBookResponseWrapper
// @Router /books/all [get] .
func (r *Route) GetAll(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = payload.GetAllBookRequest{}
		resp *presenter.ListBookResponseWrapper
	)

	temp := c.Param("unscoped")
	fmt.Println("Temp:", temp)
	if err := c.Bind(&req.Unscoped); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	fmt.Println(req)
	resp, err := r.UseCase.Book.GetAll(ctx, &req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
