package grpc

import (
	"git.teqnological.asia/teq-go/teq-echo/proto"
	"git.teqnological.asia/teq-go/teq-echo/usecase"
)

type TeqService struct {
	proto.UnimplementedTeqServiceServer
	UseCase *usecase.UseCase
}
