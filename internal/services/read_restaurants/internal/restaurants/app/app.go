package app

import (
	fxapp "github.com/Leon180/go-event-driven-microservices/internal/pkg/fxapp"
	restaurantsfx "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/fx"
)

type RestaurantsFxApp struct{}

func NewRestaurantsFxApp() *RestaurantsFxApp {
	return &RestaurantsFxApp{}
}

func (a *RestaurantsFxApp) Run() {
	app := fxapp.NewFxApp()
	app.AppendFxOptions(restaurantsfx.RestaurantsConfiguratorModule)
	app.GetLogger().Info("Starting restaurants service")
	app.Run()
}
