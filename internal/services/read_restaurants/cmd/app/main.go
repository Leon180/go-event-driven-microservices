package main

import (
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/app"

	_ "github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/docs" // swagger docs
)

//go:generate swag init -pdl 3 -o ../../docs

// @Title           Restaurants Service API
// @Version         1.0
// @Description     Simple service for restaurants resources
// @Host           localhost:8001
// @BasePath       /v1/restaurants
func main() {
	app.NewRestaurantsFxApp().Run()
}
