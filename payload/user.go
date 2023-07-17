package payload

import (
	"strings"

	"git.teqnological.asia/teq-go/teq-echo/codetype"
)

type CreateUserRequest struct {
	Name     *string `json:"name"`
	UserName *string `json:"user_name"`
	Email    *string `json:"email"`
	Gender   *string `json:"gender"`
	Team     *string `json:"team"`
	JoinDate *string `json:"join_date"`
}

type GetUserByIDRequest struct {
	ID int64 `json:"-"`
}

var orderByUser = []string{"id", "name", "created_by", "updated_by"}

type GetListUserRequest struct {
	codetype.Paginator
	SortBy    codetype.SortType `json:"sort_by,omitempty" query:"sort_by"`
	OrderBy   string            `json:"order_by,omitempty" query:"order_by"`
	Search    string            `json:"search,omitempty" query:"search"`
	CreatedBy *int64            `json:"created_by,omitempty" query:"created_by"`
}

func (g *GetListUserRequest) Format() {
	g.Paginator.Format()
	g.SortBy.Format()
	g.Search = strings.TrimSpace(g.Search)
	g.OrderBy = strings.ToLower(strings.TrimSpace(g.OrderBy))

	for i := range orderByUser {
		if g.OrderBy == orderByUser[i] {
			return
		}
	}

	g.OrderBy = ""
}

type GetAllUserRequest struct {
	Unscoped bool   `json:"unscoped"`
	Un       string `json:"un"`
}

type UpdateUserRequest struct {
	ID   int64   `json:"-"`
	Name *string `json:"name"`
}

type DeleteUserRequest struct {
	ID int64 `json:"-"`
}
