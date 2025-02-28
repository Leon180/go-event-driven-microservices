package app

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/fxapp"
	customersfx "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/fx"
)

type CustomersFxApp struct{}

func NewCustomersFxApp() *CustomersFxApp {
	return &CustomersFxApp{}
}

func (a *CustomersFxApp) Run() {
	app := fxapp.NewFxApp()
	app.AppendFxOptions(customersfx.CustomersConfiguratorModule)
	app.GetLogger().Info("Starting customers service")
	app.Run()
}
