package customizeerrors

import (
	"fmt"
	"net/http"
)

var (
	HTTPBadRequestError     = newErrorFromErrorCode(HTTPBadRequest)
	HTTPNotFoundError       = newErrorFromErrorCode(HTTPNotFound)
	HTTPInternalServerError = newErrorFromErrorCode(HTTPInternalServer)

	// File Management
	FileNotFoundError      = newErrorFromErrorCode(FileNotFound)
	DirectoryNotFoundError = newErrorFromErrorCode(DirectoryNotFound)

	// fxapp
	FxAppNotInitializedError = newErrorFromErrorCode(FxAppNotInitialized)
)

type CustomError interface {
	error
	GetStatus() int
	GetCode() int
	GetMessage() string
}

type ErrorCode int

const (
	HTTPBadRequest     ErrorCode = 400
	HTTPNotFound       ErrorCode = 404
	HTTPInternalServer ErrorCode = 500

	// File Management
	FileNotFound      ErrorCode = 10001
	DirectoryNotFound ErrorCode = 10002

	// fxapp
	FxAppNotInitialized ErrorCode = 20001
)

var errorCodeMessageMap = map[ErrorCode]string{
	HTTPBadRequest:     "bad request",
	HTTPNotFound:       "not found",
	HTTPInternalServer: "internal server error",

	// File Management
	FileNotFound:      "file not found",
	DirectoryNotFound: "directory not found",

	// fxapp
	FxAppNotInitialized: "fxapp is not initialized",
}

var errorCodeStatusMap = map[ErrorCode]int{
	HTTPBadRequest:     http.StatusBadRequest,
	HTTPNotFound:       http.StatusNotFound,
	HTTPInternalServer: http.StatusInternalServerError,

	// File Management
	FileNotFound:      http.StatusNotFound,
	DirectoryNotFound: http.StatusNotFound,

	// fxapp
	FxAppNotInitialized: http.StatusInternalServerError,
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

func NewError(status int, code int, message string) CustomError {
	return &customErrorImpl{
		status:  status,
		code:    code,
		message: message,
	}
}

func newErrorFromErrorCode(code ErrorCode) CustomError {
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
