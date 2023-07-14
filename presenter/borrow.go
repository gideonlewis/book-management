package presenter

import "git.teqnological.asia/teq-go/teq-echo/model"

type BorrowResponseWrapper struct {
	Borrow *model.Borrow `json:"borrow"`
}

type ListBorrowResponseWrapper struct {
	Borrows []model.Borrow `json:"borrows"`
	Meta    interface{}    `json:"meta"`
}
