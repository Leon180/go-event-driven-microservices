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
	InvalidIDError           = newErrorFromErrorCode(InvalidID)
	InvalidMobileNumberError = newErrorFromErrorCode(InvalidMobileNumber)
	InvalidAccountTypeError  = newErrorFromErrorCode(InvalidAccountType)
	InvalidBranchError       = newErrorFromErrorCode(InvalidBranch)
	InvalidDecimalError      = newErrorFromErrorCode(InvalidDecimal)
	InvalidEmailError        = newErrorFromErrorCode(InvalidEmail)
	InvalidNameError         = newErrorFromErrorCode(InvalidName)

	// Account
	AccountAlreadyExistsError   = newErrorFromErrorCode(AccountAlreadyExists)
	AccountNotFoundError        = newErrorFromErrorCode(AccountNotFound)
	AccountNoUpdatesError       = newErrorFromErrorCode(AccountNoUpdates)
	AccountAlreadyDeletedError  = newErrorFromErrorCode(AccountAlreadyDeleted)
	AccountAlreadyRestoredError = newErrorFromErrorCode(AccountAlreadyRestored)

	// Card
	CardAlreadyExistsError   = newErrorFromErrorCode(CardAlreadyExists)
	CardNotFoundError        = newErrorFromErrorCode(CardNotFound)
	CardNoUpdatesError       = newErrorFromErrorCode(CardNoUpdates)
	CardAlreadyDeletedError  = newErrorFromErrorCode(CardAlreadyDeleted)
	CardAlreadyRestoredError = newErrorFromErrorCode(CardAlreadyRestored)

	// Customer
	CustomerAlreadyExistsError     = newErrorFromErrorCode(CustomerAlreadyExists)
	CustomerExistsButInactiveError = newErrorFromErrorCode(CustomerExistsButInactive)
	CustomerNotFoundError          = newErrorFromErrorCode(CustomerNotFound)
	CustomerAlreadyDeletedError    = newErrorFromErrorCode(CustomerAlreadyDeleted)
	CustomerNoUpdatesError         = newErrorFromErrorCode(CustomerNoUpdates)

	// Loan
	LoanTermInvalidError    = newErrorFromErrorCode(LoanTermInvalid)
	LoanAlreadyExistsError  = newErrorFromErrorCode(LoanAlreadyExists)
	LoanNotFoundError       = newErrorFromErrorCode(LoanNotFound)
	LoanNoUpdatesError      = newErrorFromErrorCode(LoanNoUpdates)
	LoanAlreadyDeletedError = newErrorFromErrorCode(LoanAlreadyDeleted)
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
	InvalidID           ErrorCode = 21001
	InvalidMobileNumber ErrorCode = 21002
	InvalidAccountType  ErrorCode = 21003
	InvalidBranch       ErrorCode = 21004
	InvalidDecimal      ErrorCode = 21005
	InvalidEmail        ErrorCode = 21006
	InvalidName         ErrorCode = 21007

	// Account
	AccountAlreadyExists   ErrorCode = 30001
	AccountNotFound        ErrorCode = 30002
	AccountNoUpdates       ErrorCode = 30003
	AccountAlreadyDeleted  ErrorCode = 30004
	AccountAlreadyRestored ErrorCode = 30005

	// Card
	CardAlreadyExists   ErrorCode = 31001
	CardNotFound        ErrorCode = 31002
	CardNoUpdates       ErrorCode = 31003
	CardAlreadyDeleted  ErrorCode = 31004
	CardAlreadyRestored ErrorCode = 31005

	// Customer
	CustomerAlreadyExists     ErrorCode = 32001
	CustomerExistsButInactive ErrorCode = 32002
	CustomerNotFound          ErrorCode = 32003
	CustomerAlreadyDeleted    ErrorCode = 32004
	CustomerNoUpdates         ErrorCode = 32005

	// Loan
	LoanTermInvalid    ErrorCode = 33001
	LoanAlreadyExists  ErrorCode = 33002
	LoanNotFound       ErrorCode = 33003
	LoanNoUpdates      ErrorCode = 33004
	LoanAlreadyDeleted ErrorCode = 33005
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
	InvalidID:           "id is invalid",
	InvalidMobileNumber: "mobile number is invalid",
	InvalidAccountType:  "account type is invalid",
	InvalidBranch:       "branch is invalid",
	InvalidDecimal:      "decimal is invalid",
	InvalidEmail:        "email is invalid",
	InvalidName:         "name is invalid",

	// Account
	AccountAlreadyExists:   "account already exists",
	AccountNotFound:        "account not found",
	AccountNoUpdates:       "account no updates",
	AccountAlreadyDeleted:  "account already deleted",
	AccountAlreadyRestored: "account already restored",

	// Card
	CardAlreadyExists:   "card already exists",
	CardNotFound:        "card not found",
	CardNoUpdates:       "card no updates",
	CardAlreadyDeleted:  "card already deleted",
	CardAlreadyRestored: "card already restored",

	// Customer
	CustomerAlreadyExists:     "customer already exists",
	CustomerExistsButInactive: "customer exists but is inactive",
	CustomerNotFound:          "customer not found",
	CustomerAlreadyDeleted:    "customer already deleted",
	CustomerNoUpdates:         "customer no updates",

	// Loan
	LoanTermInvalid:    "loan term is invalid",
	LoanAlreadyExists:  "loan already exists",
	LoanNotFound:       "loan not found",
	LoanNoUpdates:      "loan no updates",
	LoanAlreadyDeleted: "loan already deleted",
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

	// common
	InvalidID:           http.StatusBadRequest,
	InvalidMobileNumber: http.StatusBadRequest,
	InvalidAccountType:  http.StatusBadRequest,
	InvalidBranch:       http.StatusBadRequest,
	InvalidDecimal:      http.StatusBadRequest,
	InvalidEmail:        http.StatusBadRequest,
	InvalidName:         http.StatusBadRequest,

	// Account
	AccountAlreadyExists:   http.StatusConflict,
	AccountNotFound:        http.StatusNotFound,
	AccountNoUpdates:       http.StatusOK,
	AccountAlreadyDeleted:  http.StatusOK,
	AccountAlreadyRestored: http.StatusOK,

	// Card
	CardAlreadyExists:   http.StatusConflict,
	CardNotFound:        http.StatusNotFound,
	CardNoUpdates:       http.StatusOK,
	CardAlreadyDeleted:  http.StatusOK,
	CardAlreadyRestored: http.StatusOK,

	// Customer
	CustomerAlreadyExists:     http.StatusConflict,
	CustomerExistsButInactive: http.StatusOK,
	CustomerNotFound:          http.StatusNotFound,
	CustomerAlreadyDeleted:    http.StatusOK,
	CustomerNoUpdates:         http.StatusOK,

	// Loan
	LoanTermInvalid:    http.StatusBadRequest,
	LoanAlreadyExists:  http.StatusConflict,
	LoanNotFound:       http.StatusNotFound,
	LoanNoUpdates:      http.StatusOK,
	LoanAlreadyDeleted: http.StatusOK,
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
