package usecase

import (
	"git.teqnological.asia/teq-go/teq-echo/repository"
	"git.teqnological.asia/teq-go/teq-echo/usecase/example"
	"git.teqnological.asia/teq-go/teq-echo/usecase/grpc"
	myS3 "git.teqnological.asia/teq-go/teq-echo/util/s3"
	mySES "git.teqnological.asia/teq-go/teq-echo/util/ses"
)

type UseCase struct {
	Example example.IUseCase
	GRPC    grpc.IUseCase

	SES mySES.ISES
	S3  myS3.IS3
}

func New(repo *repository.Repository) *UseCase {
	var (
		ses = mySES.NewSES()
		s3  = myS3.NewS3()
	)

	return &UseCase{
		Example: example.New(repo, ses),
		GRPC:    grpc.New(repo),
		SES:     ses,
		S3:      s3,
	}
}
