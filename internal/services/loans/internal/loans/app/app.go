package app

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/fxapp"
	loansfx "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/fx"
)

type LoansFxApp struct{}

func NewLoansFxApp() *LoansFxApp {
	return &LoansFxApp{}
}

func (a *LoansFxApp) Run() {
	app := fxapp.NewFxApp()
	app.AppendFxOptions(loansfx.LoansConfiguratorModule)
	app.GetLogger().Info("Starting loans service")
	app.Run()
}
