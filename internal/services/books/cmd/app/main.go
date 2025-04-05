package main

import (
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/app"

	_ "github.com/Leon180/go-event-driven-microservices/internal/services/books/docs" // swagger docs
)

//go:generate swag init -pdl 3 -o ../../docs

// @Title           Books Service API
// @Version         1.0
// @Description     Simple service for books resources
// @Host           localhost:8003
// @BasePath       /v1/books
func main() {
	app.NewBooksFxApp().Run()
}
