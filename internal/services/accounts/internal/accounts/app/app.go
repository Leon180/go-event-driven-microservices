package app

import (
	fxapp "github.com/Leon180/go-event-driven-microservices/internal/pkg/fxapp"
	accountsfx "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/fx"
)

type AccountFxApp struct{}

func NewAccountFxApp() *AccountFxApp {
	return &AccountFxApp{}
}

func (a *AccountFxApp) Run() {
	app := fxapp.NewFxApp()
	app.AppendFxOptions(accountsfx.AccountsConfiguratorModule)
	app.GetLogger().Info("Starting accounts service")
	app.Run()
}
