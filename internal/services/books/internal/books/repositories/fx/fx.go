package repositoriesfx

import (
	repositoriespostgres "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/repositories/postgres"
	"go.uber.org/fx"
)

// ProvideModule is the module for the repositories
// It provides the repositories:
// - repositories.SearchBooksFullInfo
// - repositories.ReadBooks
// - repositories.UpdateBooks
// - repositories.UpdateBooksWithTransaction
// dependencies:
// - *gorm.DB
// - contextloggers.ContextLogger
var ProvideModule = fx.Module(
	"booksRepositoriesProvideFx",
	fx.Provide(
		repositoriespostgres.NewSearchBooksFullInfo,
		repositoriespostgres.NewReadBooks,
		repositoriespostgres.NewUpdateBooks,
		repositoriespostgres.NewUpdateBooksWithTransaction,
	),
)
