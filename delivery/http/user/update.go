package user

import (
	"strconv"

	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"

	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
)

// Update User by id
// @Summary Update an User
// @Description Update User by id
// @Tags User
// @Accept json
// @Produce json
// @Security AuthToken
// @Param id path int true "id"
// @Param req body payload.UpdateUserRequest true "User info"
// @Success 200 {object} presenter.UserResponseWrapper
// @Router /Users/{id} [put] .
func (r *Route) Update(c echo.Context) error {
	var (
		ctx   = &teq.CustomEchoContext{Context: c}
		idStr = c.Param("id")
		resp  *presenter.UserResponseWrapper
	)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	req := payload.UpdateUserRequest{
		ID: id,
	}

	if err = c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	resp, err = r.UseCase.User.Update(ctx, &req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
