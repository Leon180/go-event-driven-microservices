package accountnumberutilities

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"

	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
)

//go:generate mockgen -source=card_number.go -destination=./mocks/card_number_mock.go -package=mocks

type CardNumberGenerator interface {
	GenerateCardNumber() string
}

func NewCardNumberGenerator() CardNumberGenerator {
	return &cardNumberGeneratorImpl{
		digits: enums.CardNumberDigits,
	}
}

type cardNumberGeneratorImpl struct {
	digits int
}

func (a *cardNumberGeneratorImpl) GenerateCardNumber() string {
	b := make([]byte, a.digits)

	rand.Read(b)

	num := binary.BigEndian.Uint64(b)%9000000000000000 + 1000000000000000

	return fmt.Sprintf("%d", num)
}
