package main

import (
	_ "github.com/Leon180/go-event-driven-microservices/internal/services/customers/docs" // swagger docs
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/app"
)

//go:generate swag init -pdl 3 -o ../../docs

// @Title           Customers Service API
// @Version         1.0
// @Description     Simple service for customers resources
// @Host           localhost:7003
// @BasePath       /v1/customers
func main() {
	app.NewCustomersFxApp().Run()
}
