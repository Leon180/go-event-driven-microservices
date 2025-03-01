package accountnumberutilities

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"

	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
)

//go:generate mockgen -source=loan_number.go -destination=./mocks/loan_number_mock.go -package=mocks

type LoanNumberGenerator interface {
	GenerateLoanNumber() string
}

func NewLoanNumberGenerator() LoanNumberGenerator {
	return &loanNumberGeneratorImpl{
		digits: enums.LoanNumberDigits,
	}
}

type loanNumberGeneratorImpl struct {
	digits int
}

func (a *loanNumberGeneratorImpl) GenerateLoanNumber() string {
	b := make([]byte, a.digits)

	rand.Read(b)

	num := binary.BigEndian.Uint64(b)%9000000000000000 + 1000000000000000

	return fmt.Sprintf("%d", num)
}
