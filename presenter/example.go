package presenter

import "git.teqnological.asia/teq-go/teq-echo/model"

type ExampleResponseWrapper struct {
	Example *model.Example `json:"example"`
}

type ListExampleResponseWrapper struct {
	Examples []model.Example `json:"examples"`
	Meta     interface{}     `json:"meta"`
}
