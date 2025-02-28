package enums

type LoanType string

const (
	LoanTypeInvalid LoanType = "invalid"
	LoanTypeHome    LoanType = "home"
	LoanTypeCar     LoanType = "car"
)

func (l LoanType) String() string {
	return string(l)
}

func (l LoanType) IsValid() bool {
	active, ok := LoanTypesActiveMap[l]
	return ok && active
}

var LoanTypesActiveMap = map[LoanType]bool{
	LoanTypeHome: true,
	LoanTypeCar:  true,
}

func (l LoanType) ToLoanTypeCode() LoanTypeCode {
	loanTypeCode, ok := LoanTypeToLoanTypeCodeMap[l]
	if !ok {
		return LoanTypeCodeInvalid
	}
	return loanTypeCode
}

var LoanTypeToLoanTypeCodeMap = map[LoanType]LoanTypeCode{
	LoanTypeHome: LoanTypeCodeHome,
	LoanTypeCar:  LoanTypeCodeCar,
}

type LoanTypeCode int

const (
	LoanTypeCodeInvalid LoanTypeCode = iota - 1
	LoanTypeCodeHome
	LoanTypeCodeCar
)

func (l LoanTypeCode) ToLoanType() LoanType {
	loanType, ok := LoanTypeCodeToLoanTypeMap[l]
	if !ok {
		return LoanTypeInvalid
	}
	return loanType
}

var LoanTypeCodeToLoanTypeMap = map[LoanTypeCode]LoanType{
	LoanTypeCodeHome: LoanTypeHome,
	LoanTypeCodeCar:  LoanTypeCar,
}
