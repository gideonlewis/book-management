package user

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/util/myerror"
)

func (u *UseCase) Delete(ctx context.Context, req *payload.DeleteUserRequest) error {
	myUser, err := u.UserRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return myerror.ErrExampleNotFound()
		}

		return myerror.ErrExampleGet(err)
	}

	err = u.UserRepo.Delete(ctx, myUser, false)
	if err != nil {
		return myerror.ErrExampleDelete(err)
	}

	return nil
}
