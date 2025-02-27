package fxapp

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	"go.uber.org/fx"
)

type FxApp interface {
	Run()
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Wait() <-chan fx.ShutdownSignal

	AppendFxOptions(options ...fx.Option)
	GetOptions() []fx.Option
	GetLogger() loggers.Logger
	GetEnvironment() enums.Environment
}
