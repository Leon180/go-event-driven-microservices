package app

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/fxapp"
	cardsfx "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/fx"
)

type CardsFxApp struct{}

func NewCardsFxApp() *CardsFxApp {
	return &CardsFxApp{}
}

func (a *CardsFxApp) Run() {
	app := fxapp.NewFxApp()
	app.AppendFxOptions(cardsfx.CardsConfiguratorModule)
	app.GetLogger().Info("Starting cards service")
	app.Run()
}
