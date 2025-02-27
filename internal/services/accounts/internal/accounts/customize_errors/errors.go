package customizeerrors

import (
	"fmt"
	"net/http"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
)

var (
	HTTPBadRequestError     = newErrorFromErrorCode(HTTPBadRequest)
	HTTPNotFoundError       = newErrorFromErrorCode(HTTPNotFound)
	HTTPInternalServerError = newErrorFromErrorCode(HTTPInternalServer)

	// Account
	AccountAlreadyExistsError = newErrorFromErrorCode(AccountAlreadyExists)
	AccountNotFoundError      = newErrorFromErrorCode(AccountNotFound)
	AccountNoUpdatesError     = newErrorFromErrorCode(AccountNoUpdates)
)

type ErrorCode int

const (
	HTTPBadRequest     ErrorCode = 400
	HTTPNotFound       ErrorCode = 404
	HTTPInternalServer ErrorCode = 500

	// Account
	AccountAlreadyExists ErrorCode = 1001
	AccountNotFound      ErrorCode = 1002
	AccountNoUpdates     ErrorCode = 1003
)

var errorCodeMessageMap = map[ErrorCode]string{
	HTTPBadRequest:     "bad request",
	HTTPNotFound:       "not found",
	HTTPInternalServer: "internal server error",

	// Account
	AccountAlreadyExists: "account already exists",
	AccountNotFound:      "account not found",
	AccountNoUpdates:     "account no updates",
}

var errorCodeStatusMap = map[ErrorCode]int{
	HTTPBadRequest:     http.StatusBadRequest,
	HTTPNotFound:       http.StatusNotFound,
	HTTPInternalServer: http.StatusInternalServerError,

	// Account
	AccountAlreadyExists: http.StatusConflict,
	AccountNotFound:      http.StatusNotFound,
	AccountNoUpdates:     http.StatusOK,
}

func (e ErrorCode) GetCode() int {
	return int(e)
}

func (e ErrorCode) GetMessage() string {
	return errorCodeMessageMap[e]
}

func (e ErrorCode) GetStatus() int {
	return errorCodeStatusMap[e]
}

func NewError(status int, code int, message string) customizeerrors.CustomError {
	return &customErrorImpl{
		status:  status,
		code:    code,
		message: message,
	}
}

func newErrorFromErrorCode(code ErrorCode) customizeerrors.CustomError {
	return NewError(code.GetStatus(), code.GetCode(), code.GetMessage())
}

type customErrorImpl struct {
	status  int
	code    int
	message string
}

func (e *customErrorImpl) GetStatus() int {
	return e.status
}

func (e *customErrorImpl) GetCode() int {
	return e.code
}

func (e *customErrorImpl) GetMessage() string {
	return e.message
}

func (e *customErrorImpl) Error() string {
	return fmt.Sprintf("status: %d, code: %d, message: %s", e.status, e.code, e.message)
}
