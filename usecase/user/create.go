package user

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"git.teqnological.asia/teq-go/teq-echo/model"
	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
	"git.teqnological.asia/teq-go/teq-echo/util/myerror"
	"gorm.io/gorm"
)

const DAY_STANDARD = "2006-01-02T15:04:05Z07:00"

func (u *UseCase) validateCreate(ctx context.Context, req *payload.CreateUserRequest) error {
	if req.Name == nil {
		return myerror.ErrUserInvalidParam("name")
	}

	*req.Name = strings.TrimSpace(*req.Name)
	if len(*req.Name) == 0 {
		return myerror.ErrUserInvalidParam("name")
	}

	// check email format
	if !checkEmailDomain(*req.Email) {
		return myerror.ErrUserInvalidParam("Email must be a valid domain: @teqnological.asia | @gmail.com")
	}

	// check username && email existed
	users, err := u.UserRepo.GetAll(ctx, true)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}

		return myerror.ErrUserGet(err)
	}

	for _, user := range users {
		if user.UserName == *req.UserName {
			return myerror.ErrUserInvalidParam("user_name already existed")
		}

		if user.Email == *req.Email {
			return myerror.ErrUserInvalidParam("email already exists")
		}
	}

	return nil
}

func (u *UseCase) Create(
	ctx context.Context,
	req *payload.CreateUserRequest,
) (*presenter.UserResponseWrapper, error) {
	if err := u.validateCreate(ctx, req); err != nil {
		return nil, err
	}

	joinDate, err := checkJoinDate(req.JoinDate)
	if err != nil {
		return nil, myerror.ErrUserCreate(err)
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
		return nil, myerror.ErrUserCreate(err)
	}

	return &presenter.UserResponseWrapper{User: myUser}, nil
}

func checkJoinDate(joinDate *string) (time.Time, error) {
	if joinDate == nil {
		return time.Parse(DAY_STANDARD, time.Now().UTC().Format(time.RFC3339))
	}

	return time.Parse(DAY_STANDARD, *joinDate)
}

func checkEmailDomain(email string) bool {

	fmt.Println("Checking email domain")
	pattern := `^[a-zA-Z0-9_.+-]+@(teqnological\.asia|gmail\.com)$`
	match, _ := regexp.MatchString(pattern, email)

	return match
}
