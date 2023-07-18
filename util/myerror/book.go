package myerror

import (
	"fmt"
	"net/http"

	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
)

func ErrBookGet(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10000",
		Message:   "Failed to get book.",
		IsSentry:  true,
	}
}

func ErrBookCreate(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10001",
		Message:   "Failed to create book.",
		IsSentry:  true,
	}
}

func ErrBookUpdate(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10002",
		Message:   "Failed to update book.",
		IsSentry:  true,
	}
}

func ErrBookDelete(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10003",
		Message:   "Failed to delete book.",
		IsSentry:  true,
	}
}

func ErrBookNotFound() teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusNotFound,
		ErrorCode: "10004",
		Message:   "Book not found.",
		IsSentry:  false,
	}
}

func ErrBookInvalidParam(param string) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "10005",
		Message:   fmt.Sprintf("Invalid paramemter: `%s`.", param),
		IsSentry:  false,
	}
}
