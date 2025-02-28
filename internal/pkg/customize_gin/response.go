package customizegin

import (
	"net/http"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/gin-gonic/gin"
)

type JSONResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   *APIError   `json:"error,omitempty"`
}

type APIError struct {
	Status  int    `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ResponseSuccess(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, JSONResponse{
		Success: true,
		Data:    data,
		Message: message,
	})
}

func ResponseError(c *gin.Context, data interface{}, message string, err error) {
	var (
		status       int
		code         int
		errorMessage string
	)
	if customErr, ok := err.(customizeerrors.CustomError); ok {
		status = customErr.GetStatus()
		code = customErr.GetCode()
		errorMessage = customErr.GetMessage()
	} else {
		status = http.StatusInternalServerError
		code = http.StatusInternalServerError
		errorMessage = err.Error()
	}
	c.JSON(status, JSONResponse{
		Success: false,
		Data:    data,
		Message: message,
		Error: &APIError{
			Status:  status,
			Code:    code,
			Message: errorMessage,
		},
	})
}
