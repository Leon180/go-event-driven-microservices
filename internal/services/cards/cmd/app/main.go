package main

import (
	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/app"

	_ "github.com/Leon180/go-event-driven-microservices/internal/services/cards/docs" // swagger docs
)

//go:generate swag init -pdl 3 -o ../../docs

// @Title           Cards Service API
// @Version         1.0
// @Description     Simple service for cards resources
// @Host           localhost:7002
// @BasePath       /v1/cards
func main() {
	app.NewCardsFxApp().Run()
}
