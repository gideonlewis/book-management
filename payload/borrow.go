package payload

import (
	"strings"

	"git.teqnological.asia/teq-go/teq-echo/codetype"
)

type CreateBorrowRequest struct {
	Name *string `json:"name"`
}

type GetBorrowByIDRequest struct {
	ID int64 `json:"-"`
}

var orderByBorrow = []string{"id", "name", "created_by", "updated_by"}

type GetListBorrowRequest struct {
	codetype.Paginator
	SortBy    codetype.SortType `json:"sort_by,omitempty" query:"sort_by"`
	OrderBy   string            `json:"order_by,omitempty" query:"order_by"`
	Search    string            `json:"search,omitempty" query:"search"`
	CreatedBy *int64            `json:"created_by,omitempty" query:"created_by"`
}

func (g *GetListBorrowRequest) Format() {
	g.Paginator.Format()
	g.SortBy.Format()
	g.Search = strings.TrimSpace(g.Search)
	g.OrderBy = strings.ToLower(strings.TrimSpace(g.OrderBy))

	for i := range orderByBorrow {
		if g.OrderBy == orderByBorrow[i] {
			return
		}
	}

	g.OrderBy = ""
}

type GetAllBorrowRequest struct {
	Unscoped bool `json:"unscoped"`
}

type UpdateBorrowRequest struct {
	ID   int64   `json:"-"`
	Name *string `json:"name"`
}

type DeleteBorrowRequest struct {
	ID int64 `json:"-"`
}
