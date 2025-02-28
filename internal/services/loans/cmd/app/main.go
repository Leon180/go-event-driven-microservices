package main

import (
	_ "github.com/Leon180/go-event-driven-microservices/internal/services/loans/docs" // swagger docs
	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/app"
)

//go:generate swag init -pdl 3 -o ../../docs

// @Title           Loans Service API
// @Version         1.0
// @Description     Simple service for loans resources
// @Host           localhost:7003
// @BasePath       /v1/loans
func main() {
	app.NewLoansFxApp().Run()
}
