package book

import (
	"strconv"

	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"

	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
)

// Update Book by id
// @Summary Update an Book
// @Description Update Book by id
// @Tags Book
// @Accept json
// @Produce json
// @Security AuthToken
// @Param id path int true "id"
// @Param req body payload.UpdateBookRequest true "Book info"
// @Success 200 {object} presenter.BookResponseWrapper
// @Router /Books/{id} [put] .
func (r *Route) Update(c echo.Context) error {
	var (
		ctx   = &teq.CustomEchoContext{Context: c}
		idStr = c.Param("id")
		resp  *presenter.BookResponseWrapper
	)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	req := payload.UpdateBookRequest{
		ID: id,
	}

	if err = c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	resp, err = r.UseCase.Book.Update(ctx, &req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
