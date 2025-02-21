package main

import "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/app"

func main() {
	app.NewAccountFxApp().Run()
}
