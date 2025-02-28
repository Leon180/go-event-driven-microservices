package fxcustomizelogger

import (
	"strings"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

// Ref: https://articles.wesionary.team/logging-interfaces-in-go-182c28be3d18

var FxLogger = fx.WithLogger(func(logger loggers.Logger) fxevent.Logger {
	return NewCustomFxLogger(logger)
})

func NewCustomFxLogger(logger loggers.Logger) fxevent.Logger {
	return &fxCustomLogger{Logger: logger}
}

type fxCustomLogger struct {
	loggers.Logger
}

// LogEvent logs the given event to the provided Zap logger.
func (l *fxCustomLogger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.Debugw("OnStart hook executing", loggers.Fields{"caller": e.CallerName, "function": e.FunctionName})
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.Errorw("OnStart hook failed", loggers.Fields{"caller": e.CallerName, "callee": e.CallerName, "error": e.Err})
		} else {
			l.Debugw("OnStart hook executed", loggers.Fields{"caller": e.CallerName, "callee": e.FunctionName, "runtime": e.Runtime.String()})
		}
	case *fxevent.OnStopExecuting:
		l.Debugw("OnStop hook executing", loggers.Fields{"callee": e.FunctionName, "caller": e.CallerName})
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.Errorw("OnStop hook failed", loggers.Fields{"caller": e.CallerName, "callee": e.CallerName, "error": e.Err})
		} else {
			l.Debugw("OnStop hook executed", loggers.Fields{"caller": e.CallerName, "callee": e.FunctionName, "runtime": e.Runtime.String()})
		}
	case *fxevent.Supplied:
		if e.Err != nil {
			l.Errorw("error encountered while applying options", loggers.Fields{"type": e.TypeName, "stacktrace": e.StackTrace, "module": e.ModuleName, "error": e.Err})
		} else {
			l.Debugw("supplied", loggers.Fields{"type": e.TypeName, "stacktrace": e.StackTrace, "module": e.ModuleName})
		}
	case *fxevent.Provided:
		for _, rtype := range e.OutputTypeNames {
			l.Debugw("provided", loggers.Fields{"constructor": e.ConstructorName, "stacktrace": e.StackTrace, "module": e.ModuleName, "type": rtype, "private": e.Private})
		}
		if e.Err != nil {
			l.Errorw("error encountered while applying options", loggers.Fields{"module": e.ModuleName, "stacktrace": e.StackTrace, "error": e.Err})
		}
	case *fxevent.Replaced:
		for _, rtype := range e.OutputTypeNames {
			l.Debugw("replaced", loggers.Fields{"stacktrace": e.StackTrace, "module": e.ModuleName, "type": rtype})
		}
		if e.Err != nil {
			l.Errorw("error encountered while replacing", loggers.Fields{"module": e.ModuleName, "stacktrace": e.StackTrace, "error": e.Err})
		}
	case *fxevent.Decorated:
		for _, rtype := range e.OutputTypeNames {
			l.Debugw("decorated", loggers.Fields{"decorator": e.DecoratorName, "stacktrace": e.StackTrace, "module": e.ModuleName, "type": rtype})
		}
		if e.Err != nil {
			l.Errorw("error encountered while applying options", loggers.Fields{"module": e.ModuleName, "stacktrace": e.StackTrace, "error": e.Err})
		}
	case *fxevent.Run:
		if e.Err != nil {
			l.Errorw("error returned", loggers.Fields{"module": e.ModuleName, "name": e.Name, "kind": e.Kind, "error": e.Err})
		} else {
			l.Debugw("run", loggers.Fields{"module": e.ModuleName, "name": e.Name, "kind": e.Kind})
		}
	case *fxevent.Invoking:
		// Do not log stack as it will make logs hard to read.
		l.Debugw("invoking", loggers.Fields{"module": e.ModuleName, "function": e.FunctionName})
	case *fxevent.Invoked:
		if e.Err != nil {
			l.Errorw("invoke failed", loggers.Fields{"error": e.Err, "stack": e.Trace, "function": e.FunctionName, "module": e.ModuleName})
		}
	case *fxevent.Stopping:
		l.Debugw("received signal", loggers.Fields{"signal": strings.ToUpper(e.Signal.String())})
	case *fxevent.Stopped:
		if e.Err != nil {
			l.Errorw("stop failed", loggers.Fields{"error": e.Err})
		}
	case *fxevent.RollingBack:
		l.Errorw("start failed, rolling back", loggers.Fields{"error": e.StartErr})
	case *fxevent.RolledBack:
		if e.Err != nil {
			l.Errorw("rollback failed", loggers.Fields{"error": e.Err})
		}
	case *fxevent.Started:
		if e.Err != nil {
			l.Errorw("start failed", loggers.Fields{"error": e.Err})
		} else {
			l.Debug("started")
		}
	case *fxevent.LoggerInitialized:
		if e.Err != nil {
			l.Errorw("custom logger initialization failed", loggers.Fields{"error": e.Err})
		} else {
			l.Debugw("initialized custom fxevent.logger", loggers.Fields{"function": e.ConstructorName})
		}
	}
}
