package presenter

import "git.teqnological.asia/teq-go/teq-echo/model"

type BookResponseWrapper struct {
	Book *model.Book `json:"book"`
}

type ListBookResponseWrapper struct {
	Books []model.Book `json:"books"`
	Meta  interface{}  `json:"meta"`
}
