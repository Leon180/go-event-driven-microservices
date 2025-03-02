package main

import (
	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/app"

	_ "github.com/Leon180/go-event-driven-microservices/internal/services/loans/docs" // swagger docs
)

//go:generate swag init -pdl 3 -o ../../docs

// @Title           Loans Service API
// @Version         1.0
// @Description     Simple service for loans resources
// @Host           localhost:7004
// @BasePath       /v1/loans
func main() {
	app.NewLoansFxApp().Run()
}
