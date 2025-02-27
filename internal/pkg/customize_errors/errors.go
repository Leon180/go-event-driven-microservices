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

	// common
	IDInvalidError = newErrorFromErrorCode(IDInvalid)

	// Account
	AccountAlreadyExistsError = newErrorFromErrorCode(AccountAlreadyExists)
	AccountNotFoundError      = newErrorFromErrorCode(AccountNotFound)
	AccountNoUpdatesError     = newErrorFromErrorCode(AccountNoUpdates)

	AccountInvalidMobileNumberError = newErrorFromErrorCode(AccountInvalidMobileNumber)
	AccountInvalidAccountTypeError  = newErrorFromErrorCode(AccountInvalidAccountType)
	AccountInvalidBranchError       = newErrorFromErrorCode(AccountInvalidBranch)
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

	// common
	IDInvalid ErrorCode = 21001

	// Account
	AccountAlreadyExists ErrorCode = 30001
	AccountNotFound      ErrorCode = 30002
	AccountNoUpdates     ErrorCode = 30003

	AccountInvalidMobileNumber ErrorCode = 30004
	AccountInvalidAccountType  ErrorCode = 30005
	AccountInvalidBranch       ErrorCode = 30006
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

	// common
	IDInvalid: "id is invalid",

	// Account
	AccountAlreadyExists: "account already exists",
	AccountNotFound:      "account not found",
	AccountNoUpdates:     "account no updates",

	AccountInvalidMobileNumber: "invalid mobile number",
	AccountInvalidAccountType:  "invalid account type",
	AccountInvalidBranch:       "invalid branch",
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

	// Account
	AccountAlreadyExists: http.StatusConflict,
	AccountNotFound:      http.StatusNotFound,
	AccountNoUpdates:     http.StatusOK,

	AccountInvalidMobileNumber: http.StatusBadRequest,
	AccountInvalidAccountType:  http.StatusBadRequest,
	AccountInvalidBranch:       http.StatusBadRequest,
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
