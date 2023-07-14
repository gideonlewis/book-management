package book

import (
	"strconv"

	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"

	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
)

// GetByID example by id
// @Summary Get an example
// @Description Get example by id
// @Tags Example
// @Accept json
// @Produce json
// @Security AuthToken
// @Param id path int true "id"
// @Success 200 {object} presenter.ExampleResponseWrapper
// @Router /examples/{id} [get] .
func (r *Route) GetByID(c echo.Context) error {
	var (
		ctx   = &teq.CustomEchoContext{Context: c}
		idStr = c.Param("id")
		resp  *presenter.UserResponseWrapper
	)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	resp, err = r.UseCase.Book.GetByID(ctx, &payload.GetUserByIDRequest{ID: id})
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
