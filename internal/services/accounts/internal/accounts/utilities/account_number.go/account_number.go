package accountnumberutilities

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"

	enumsaccounts "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums/accounts"
)

type AccountNumberGenerator interface {
	GenerateAccountNumber() string
}

func NewAccountNumberGenerator() AccountNumberGenerator {
	return &accountNumberGeneratorImpl{
		digits: enumsaccounts.AccountNumberDigits,
	}
}

type accountNumberGeneratorImpl struct {
	digits int
}

func (a *accountNumberGeneratorImpl) GenerateAccountNumber() string {
	b := make([]byte, a.digits)

	rand.Read(b)

	num := binary.BigEndian.Uint64(b)%900000000000 + 100000000000

	return fmt.Sprintf("%d", num)
}
