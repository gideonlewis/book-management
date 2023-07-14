package presenter

import "git.teqnological.asia/teq-go/teq-echo/model"

type UserResponseWrapper struct {
	User *model.User `json:"user"`
}

type ListUserResponseWrapper struct {
	Users []model.User `json:"users"`
	Meta  interface{}  `json:"meta"`
}
