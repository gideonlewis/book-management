package payload

import (
	"strings"

	"git.teqnological.asia/teq-go/teq-echo/codetype"
)

type StatisticBorrowRequest struct {
	Unscoped bool    `json:"unscoped"`
	From     *string `json:"from,omitempty" query:"from"`
	To       *string `json:"to,omitempty" query:"to"`
}
type CreateBorrowRequest struct {
	UserID   *int64 `json:"user_id"`
	BookID   *int64 `json:"book_id"`
	Quantity *int64 `json:"quantity"`
}

type GetBorrowByIDRequest struct {
	ID int64 `json:"-"`
}

var orderByBorrow = []string{"id", "title", "user_id", "created_by", "updated_by"}

type GetListBorrowRequest struct {
	codetype.Paginator
	SortBy    codetype.SortType `json:"sort_by,omitempty" query:"sort_by"`
	OrderBy   string            `json:"order_by,omitempty" query:"order_by"`
	Search    string            `json:"search,omitempty" query:"search"`
	UserID    *int64            `json:"user_id,omitempty" query:"user_id"`
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
	ID         int64   `json:"-"`
	ReturnDate *string `json:"return_date"`
}

type DeleteBorrowRequest struct {
	ID int64 `json:"-"`
}
