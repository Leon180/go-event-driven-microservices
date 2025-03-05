package main

import (
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/app"

	_ "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/docs" // swagger docs
)

//go:generate swag init -pdl 3 -o ../../docs

// @Title           Accounts Service API
// @Version         1.0
// @Description     Simple service for accounts resources
// @Host           localhost:7001
// @BasePath       /v1/accounts
func main() {
	app.NewAccountFxApp().Run()
}
