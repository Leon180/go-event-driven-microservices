package errors

type CustomError interface {
	error
	GetStatus() int
	GetCode() int
	GetMessage() string
}
