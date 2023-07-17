package user

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"git.teqnological.asia/teq-go/teq-echo/client/mysql"
	"git.teqnological.asia/teq-go/teq-echo/model"
	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
	"git.teqnological.asia/teq-go/teq-echo/util/myerror"
	"gorm.io/gorm"
)

const DAY_STANDARD = "2006-01-02T15:04:05Z07:00"

func (u *UseCase) validateCreate(req *payload.CreateUserRequest) error {
	if req.Name == nil {
		return myerror.ErrExampleInvalidParam("name")
	}

	*req.Name = strings.TrimSpace(*req.Name)
	if len(*req.Name) == 0 {
		return myerror.ErrExampleInvalidParam("name")
	}

	return nil
}

func (u *UseCase) Create(
	ctx context.Context,
	req *payload.CreateUserRequest,
) (*presenter.UserResponseWrapper, error) {
	if err := u.validateCreate(req); err != nil {
		return nil, err
	}

	joinDate, err := checkJoinDate(req.JoinDate)
	if err != nil {
		return nil, myerror.ErrExampleCreate(err)
	}

	fmt.Println(checkEmail(*req.Email))
	if !checkEmail(*req.Email) {
		return &presenter.UserResponseWrapper{User: &model.User{}}, myerror.ErrExampleCreate(errors.New("invalid email"))
	}

	if isEmailExists(mysql.GetDB(), *req.Email) {
		return &presenter.UserResponseWrapper{User: &model.User{}}, myerror.ErrExampleCreate(errors.New("email already exists"))
	}

	if isUsernameExists(mysql.GetDB(), *req.UserName) {
		return &presenter.UserResponseWrapper{User: &model.User{}}, myerror.ErrExampleCreate(errors.New("user_name already exists"))
	}

	myUser := &model.User{
		Name:      *req.Name,
		UserName:  *req.UserName,
		Email:     *req.Email,
		Gender:    *req.Gender,
		Team:      *req.Team,
		JoinDate:  joinDate,
		CreatedBy: 1, // must be validate logged user.
	}

	err = u.UserRepo.Create(ctx, myUser)
	if err != nil {
		return nil, myerror.ErrExampleCreate(err)
	}

	return &presenter.UserResponseWrapper{User: myUser}, nil
}

func checkEmail(email string) bool {
	pattern := `^[a-zA-Z0-9_.+-]+@(teqnological\.asia|gmail\.com)$`
	match, _ := regexp.MatchString(pattern, email)

	return match
}

func isUsernameExists(db *gorm.DB, username string) bool {
	var count int64
	db.Table("users").Where("user_name = ?", username).Count(&count)
	return count > 0
}

func isEmailExists(db *gorm.DB, email string) bool {
	var count int64
	db.Table("users").Where("email = ?", email).Count(&count)
	return count > 0
}

func checkJoinDate(joinDate *string) (time.Time, error) {
	if joinDate == nil {
		return time.Parse(DAY_STANDARD, time.Now().UTC().Format(time.RFC3339))
	}

	return time.Parse(DAY_STANDARD, *joinDate)
}
