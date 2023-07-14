package example

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
	"git.teqnological.asia/teq-go/teq-echo/util/myerror"
)

func (u *UseCase) GetByID(ctx context.Context, req *payload.GetByIDRequest) (*presenter.ExampleResponseWrapper, error) {
	myExample, err := u.ExampleRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerror.ErrExampleNotFound()
		}

		return nil, myerror.ErrExampleGet(err)
	}

	return &presenter.ExampleResponseWrapper{Example: myExample}, nil
}
