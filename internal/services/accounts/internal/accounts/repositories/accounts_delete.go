package repositories

import (
	"context"
)

type DeleteAccountByID interface {
	DeleteAccountByID(ctx context.Context, id string) error
}
