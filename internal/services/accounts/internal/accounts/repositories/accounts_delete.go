package repositories

import (
	"context"
)

//go:generate mockgen -source=accounts_delete.go -destination=./mocks/accounts_delete_mock.go -package=mocks
type DeleteAccountByID interface {
	DeleteAccountByID(ctx context.Context, id string) error
}
