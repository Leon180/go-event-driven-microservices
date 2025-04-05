package app

import (
	fxapp "github.com/Leon180/go-event-driven-microservices/internal/pkg/fxapp"
	booksfx "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/fx"
)

type BooksFxApp struct{}

func NewBooksFxApp() *BooksFxApp {
	return &BooksFxApp{}
}

func (a *BooksFxApp) Run() {
	app := fxapp.NewFxApp()
	app.AppendFxOptions(booksfx.BooksConfiguratorModule)
	app.GetLogger().Info("Starting books service")
	app.Run()
}
