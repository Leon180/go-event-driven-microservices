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
	InvalidCityError         = newErrorFromErrorCode(InvalidCity)
	InvalidCountryError      = newErrorFromErrorCode(InvalidCountry)
	InvalidCategoryError     = newErrorFromErrorCode(InvalidCategory)

	// action
	AlreadyDeletedError = newErrorFromErrorCode(AlreadyDeleted)
	AlreadyActiveError  = newErrorFromErrorCode(AlreadyActive)
	NoChangesError      = newErrorFromErrorCode(NoChanges)

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
	InvalidLoanTypeError    = newErrorFromErrorCode(InvalidLoanType)

	// Restaurant
	RestaurantNameEmptyError                = newErrorFromErrorCode(RestaurantNameEmpty)
	RestaurantAlreadyExistsError            = newErrorFromErrorCode(RestaurantAlreadyExists)
	RestaurantNotFoundError                 = newErrorFromErrorCode(RestaurantNotFound)
	RestaurantAlreadyExistsButInactiveError = newErrorFromErrorCode(RestaurantAlreadyExistsButInactive)
	RestaurantEventCastingError             = newErrorFromErrorCode(RestaurantEventCasting)

	// Branch
	BranchNameEmptyError     = newErrorFromErrorCode(BranchNameEmpty)
	BranchAlreadyExistsError = newErrorFromErrorCode(BranchAlreadyExists)
	BranchNotFoundError      = newErrorFromErrorCode(BranchNotFound)

	// Address
	AddressEmptyError    = newErrorFromErrorCode(AddressEmpty)
	AddressNotFoundError = newErrorFromErrorCode(AddressNotFound)

	// Price Range
	PriceRangeEmptyError    = newErrorFromErrorCode(PriceRangeEmpty)
	PriceRangeNotFoundError = newErrorFromErrorCode(PriceRangeNotFound)
	PriceRangeInvalidError  = newErrorFromErrorCode(PriceRangeInvalid)

	// Category
	CategoryNameEmptyError = newErrorFromErrorCode(CategoryNameEmpty)
	CategoryIDEmptyError   = newErrorFromErrorCode(CategoryIDEmpty)

	// Table
	TableCapacityInvalidError      = newErrorFromErrorCode(TableCapacityInvalid)
	TableAvailableTimeInvalidError = newErrorFromErrorCode(TableAvailableTimeInvalid)
	TableAlreadyExistsError        = newErrorFromErrorCode(TableAlreadyExists)
	TableNotFoundError             = newErrorFromErrorCode(TableNotFound)

	// Book
	BookAmountInvalidError             = newErrorFromErrorCode(BookAmountInvalid)
	BookAlreadyExistsError             = newErrorFromErrorCode(BookAlreadyExists)
	BookNotFoundError                  = newErrorFromErrorCode(BookNotFound)
	BookTableAndAvailableNotMatchError = newErrorFromErrorCode(BookTableAndAvailableNotMatch)
	BookAlreadyExistsButInactiveError  = newErrorFromErrorCode(BookAlreadyExistsButInactive)
	BookEventCastingError              = newErrorFromErrorCode(BookEventCasting)

	// Failed Message
	FailedMessageAlreadyExistsError = newErrorFromErrorCode(FailedMessageAlreadyExists)

	// rabbitmq
	RabbitmqConnectionError = newErrorFromErrorCode(RabbitmqConnection)

	// message
	MessageTypeInvalidError = newErrorFromErrorCode(MessageTypeInvalid)
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
	InvalidCity         ErrorCode = 21008
	InvalidCountry      ErrorCode = 21009
	InvalidCategory     ErrorCode = 21010

	// action
	AlreadyDeleted ErrorCode = 22001
	AlreadyActive  ErrorCode = 22002
	NoChanges      ErrorCode = 22003

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
	InvalidLoanType    ErrorCode = 33006

	// Restaurant
	RestaurantNameEmpty                ErrorCode = 40001
	RestaurantAlreadyExists            ErrorCode = 40002
	RestaurantNotFound                 ErrorCode = 40003
	RestaurantAlreadyExistsButInactive ErrorCode = 40004
	RestaurantEventCasting             ErrorCode = 40005

	// Branch
	BranchNameEmpty     ErrorCode = 41001
	BranchAlreadyExists ErrorCode = 41002
	BranchNotFound      ErrorCode = 41003

	// Address
	AddressEmpty    ErrorCode = 42001
	AddressNotFound ErrorCode = 42002

	// Price Range
	PriceRangeEmpty    ErrorCode = 43001
	PriceRangeNotFound ErrorCode = 43002
	PriceRangeInvalid  ErrorCode = 43003

	// Category
	CategoryNameEmpty       ErrorCode = 44001
	CategoryIDEmpty         ErrorCode = 44002
	CategoryNotFound        ErrorCode = 44003
	CategoryAlreadyExists   ErrorCode = 44004
	CategoryAlreadyDeleted  ErrorCode = 44005
	CategoryAlreadyRestored ErrorCode = 44006

	// Table
	TableCapacityInvalid      ErrorCode = 45001
	TableAvailableTimeInvalid ErrorCode = 45002
	TableAlreadyExists        ErrorCode = 45003
	TableNotFound             ErrorCode = 45004

	// Book
	BookAmountInvalid             ErrorCode = 46001
	BookAlreadyExists             ErrorCode = 46002
	BookNotFound                  ErrorCode = 46003
	BookTableAndAvailableNotMatch ErrorCode = 46004
	BookAlreadyExistsButInactive  ErrorCode = 46005
	BookEventCasting              ErrorCode = 46006

	// Failed Message
	FailedMessageAlreadyExists ErrorCode = 47001

	// rabbitmq
	RabbitmqConnection ErrorCode = 50001

	// message
	MessageTypeInvalid ErrorCode = 60001
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
	InvalidCity:         "city is invalid",
	InvalidCountry:      "country is invalid",
	InvalidCategory:     "category is invalid",

	// action
	AlreadyDeleted: "already deleted",
	AlreadyActive:  "already active",
	NoChanges:      "no changes",

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
	InvalidLoanType:    "loan type is invalid",

	// Restaurant
	RestaurantNameEmpty:                "restaurant name is empty",
	RestaurantAlreadyExists:            "restaurant already exists",
	RestaurantNotFound:                 "restaurant not found",
	RestaurantAlreadyExistsButInactive: "restaurant already exists but is inactive",
	RestaurantEventCasting:             "restaurant event casting error",

	// Branch
	BranchNameEmpty:     "branch name is empty",
	BranchAlreadyExists: "branch already exists",
	BranchNotFound:      "branch not found",

	// Address
	AddressEmpty:    "address is empty",
	AddressNotFound: "address not found",

	// Price Range
	PriceRangeEmpty:    "price range is empty",
	PriceRangeNotFound: "price range not found",
	PriceRangeInvalid:  "price range is invalid",

	// Category
	CategoryNameEmpty:       "category name is empty",
	CategoryIDEmpty:         "category id is empty",
	CategoryNotFound:        "category not found",
	CategoryAlreadyExists:   "category already exists",
	CategoryAlreadyDeleted:  "category already deleted",
	CategoryAlreadyRestored: "category already restored",

	// Table
	TableCapacityInvalid:      "table capacity is invalid",
	TableAvailableTimeInvalid: "table available time is invalid",
	TableAlreadyExists:        "table already exists",
	TableNotFound:             "table not found",

	// Book
	BookAmountInvalid:             "book amount is invalid",
	BookAlreadyExists:             "book already exists",
	BookNotFound:                  "book not found",
	BookTableAndAvailableNotMatch: "book table and available not match",
	BookAlreadyExistsButInactive:  "book already exists but is inactive",
	BookEventCasting:              "book event casting error",

	// Failed Message
	FailedMessageAlreadyExists: "failed message already exists",

	// rabbitmq
	RabbitmqConnection: "rabbitmq connection error",

	// message
	MessageTypeInvalid: "message type is invalid",
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
	InvalidCity:         http.StatusBadRequest,
	InvalidCountry:      http.StatusBadRequest,
	InvalidCategory:     http.StatusBadRequest,

	// action
	AlreadyDeleted: http.StatusOK,
	AlreadyActive:  http.StatusOK,
	NoChanges:      http.StatusOK,

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
	InvalidLoanType:    http.StatusBadRequest,

	// Restaurant
	RestaurantNameEmpty:                http.StatusBadRequest,
	RestaurantAlreadyExists:            http.StatusConflict,
	RestaurantNotFound:                 http.StatusNotFound,
	RestaurantAlreadyExistsButInactive: http.StatusConflict,
	RestaurantEventCasting:             http.StatusBadRequest,

	// Branch
	BranchNameEmpty:     http.StatusBadRequest,
	BranchAlreadyExists: http.StatusConflict,
	BranchNotFound:      http.StatusNotFound,

	// Address
	AddressEmpty:    http.StatusBadRequest,
	AddressNotFound: http.StatusNotFound,

	// Price Range
	PriceRangeEmpty:    http.StatusBadRequest,
	PriceRangeNotFound: http.StatusNotFound,
	PriceRangeInvalid:  http.StatusBadRequest,

	// Category
	CategoryNameEmpty:       http.StatusBadRequest,
	CategoryIDEmpty:         http.StatusBadRequest,
	CategoryNotFound:        http.StatusNotFound,
	CategoryAlreadyExists:   http.StatusConflict,
	CategoryAlreadyDeleted:  http.StatusOK,
	CategoryAlreadyRestored: http.StatusOK,

	// Table
	TableCapacityInvalid:      http.StatusBadRequest,
	TableAvailableTimeInvalid: http.StatusBadRequest,
	TableAlreadyExists:        http.StatusConflict,
	TableNotFound:             http.StatusNotFound,

	// Book
	BookAmountInvalid:             http.StatusBadRequest,
	BookAlreadyExists:             http.StatusConflict,
	BookNotFound:                  http.StatusNotFound,
	BookTableAndAvailableNotMatch: http.StatusConflict,
	BookAlreadyExistsButInactive:  http.StatusConflict,
	BookEventCasting:              http.StatusBadRequest,

	// Failed Message
	FailedMessageAlreadyExists: http.StatusConflict,

	// rabbitmq
	RabbitmqConnection: http.StatusInternalServerError,

	// message
	MessageTypeInvalid: http.StatusBadRequest,
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
