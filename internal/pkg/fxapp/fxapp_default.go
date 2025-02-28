package fxapp

import (
	"context"
	"time"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/environments"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	fxcustomizelogger "github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers/fx_customize_logger"
	zapcustomizelogger "github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers/zap_customize_logger"
	"go.uber.org/fx"
)

func NewFxApp() FxApp {
	env := environments.InitEnv()
	logger := zapcustomizelogger.NewZapLogger(env)
	return &fxAppImpl{
		logger: logger,
		env:    env,
	}
}

type fxAppImpl struct {
	options []fx.Option
	logger  loggers.Logger
	env     enums.Environment
	fxapp   *fx.App
}

func (a *fxAppImpl) Run() {
	if a.fxapp == nil {
		a.fxapp = createFxApp(a.options, a.logger)
	}
	a.fxapp.Run()
}

func (a *fxAppImpl) Start(ctx context.Context) error {
	if a.fxapp == nil {
		a.fxapp = createFxApp(a.options, a.logger)
	}
	return a.fxapp.Start(ctx)
}

func (a *fxAppImpl) Stop(ctx context.Context) error {
	if a.fxapp == nil {
		return customizeerrors.FxAppNotInitializedError
	}
	return a.fxapp.Stop(ctx)
}

func (a *fxAppImpl) Wait() <-chan fx.ShutdownSignal {
	if a.fxapp == nil {
		return nil
	}
	return a.fxapp.Wait()
}

func (a *fxAppImpl) AppendFxOptions(options ...fx.Option) {
	a.options = append(a.options, options...)
}

func (a *fxAppImpl) GetOptions() []fx.Option {
	return a.options
}

func (a *fxAppImpl) GetLogger() loggers.Logger {
	return a.logger
}

func (a *fxAppImpl) GetEnvironment() enums.Environment {
	return a.env
}

func createFxApp(
	options []fx.Option,
	logger loggers.Logger,
) *fx.App {
	return fx.New(
		fxcustomizelogger.FxLogger,
		fx.StartTimeout(30*time.Second),
		fx.ErrorHook(newFxErrorHandler(logger)),
		fx.Module("fxapp",
			options...,
		),
	)
}

type fxErrorHandler struct {
	logger loggers.Logger
}

func newFxErrorHandler(logger loggers.Logger) *fxErrorHandler {
	return &fxErrorHandler{logger: logger}
}

func (h *fxErrorHandler) HandleError(e error) {
	h.logger.Error(e)
}
