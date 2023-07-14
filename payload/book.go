package payload

import (
	"strings"

	"git.teqnological.asia/teq-go/teq-echo/codetype"
)

type CreateBookRequest struct {
	Name   *string `json:"name"`
	Author *string `json:"author"`
	Price  *int64  `json:"price"`
}

type GetBookByIDRequest struct {
	ID int64 `json:"-"`
}

var orderByBook = []string{"id", "name", "created_by", "updated_by"}

type GetListBookRequest struct {
	codetype.Paginator
	SortBy    codetype.SortType `json:"sort_by,omitempty" query:"sort_by"`
	OrderBy   string            `json:"order_by,omitempty" query:"order_by"`
	Search    string            `json:"search,omitempty" query:"search"`
	CreatedBy *int64            `json:"created_by,omitempty" query:"created_by"`
}

func (g *GetListBookRequest) Format() {
	g.Paginator.Format()
	g.SortBy.Format()
	g.Search = strings.TrimSpace(g.Search)
	g.OrderBy = strings.ToLower(strings.TrimSpace(g.OrderBy))

	for i := range orderByBook {
		if g.OrderBy == orderByBook[i] {
			return
		}
	}

	g.OrderBy = ""
}

type GetAllBookRequest struct {
	Unscoped bool `json:"unscoped"`
}

type UpdateBookRequest struct {
	ID   int64   `json:"-"`
	Name *string `json:"name"`
}

type DeleteBookRequest struct {
	ID int64 `json:"-"`
}
