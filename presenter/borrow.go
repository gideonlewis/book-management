package presenter

import (
	"git.teqnological.asia/teq-go/teq-echo/model"
)

type Statistic struct {
	ID            int64   `json:"id"`
	Title         string  `json:"title"`
	NumOfBorrowed int64   `json:"num_of_borrowed"`
	Quantity      int64   `json:"quantity"`
	Quantum       float64 `json:"quantum"`
	Detail        []struct {
		Title    string `json:"title"`
		Quantity int64  `json:"quantity"`
	} `gorm:"-" json:"detail"`
}

type BorrowStatisticResponseWrapper struct {
	Statistics []*Statistic `json:"statistics"`
	Meta       interface{}  `json:"meta"`
}

type BorrowResponseWrapper struct {
	Borrow *model.Borrow `json:"borrow"`
}

type ListBorrowResponseWrapper struct {
	Borrows []model.Borrow `json:"borrows"`
	Meta    interface{}    `json:"meta"`
}
