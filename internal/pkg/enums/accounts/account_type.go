package enumsaccounts

type AccountType string

const (
	AccountTypeInvalid  AccountType = "invalid"
	AccountTypeChecking AccountType = "checking"
	AccountTypeSavings  AccountType = "savings"
	AccountTypeCurrency AccountType = "currency"
	AccountTypeSalary   AccountType = "salary"
	AccountTypeBusiness AccountType = "business"
)

func (a AccountType) ToString() string {
	return string(a)
}

func (a AccountType) IsValid() bool {
	active, ok := AccountTypActiveMap[a]
	return ok && active
}

var AccountTypActiveMap = map[AccountType]bool{
	AccountTypeSavings:  true,
	AccountTypeChecking: true,
	AccountTypeCurrency: true,
	AccountTypeSalary:   true,
	AccountTypeBusiness: true,
}

func (a AccountType) ToAccountTypeCode() AccountTypeCode {
	code, ok := AccountTypeToCodeMap[a]
	if !ok {
		return AccountTypeCodeInvalid
	}
	return code
}

var AccountTypeToCodeMap = map[AccountType]AccountTypeCode{
	AccountTypeSavings:  AccountTypeCodeSavings,
	AccountTypeChecking: AccountTypeCodeChecking,
	AccountTypeCurrency: AccountTypeCodeCurrency,
	AccountTypeSalary:   AccountTypeCodeSalary,
	AccountTypeBusiness: AccountTypeCodeBusiness,
}

type AccountTypeCode int

const (
	AccountTypeCodeInvalid AccountTypeCode = iota - 1
	AccountTypeCodeSavings
	AccountTypeCodeChecking
	AccountTypeCodeCurrency
	AccountTypeCodeSalary
	AccountTypeCodeBusiness
)

func (a AccountTypeCode) ToAccountType() AccountType {
	accountType, ok := AccountTypeCodeToAccountTypeMap[a]
	if !ok {
		return AccountTypeInvalid
	}
	return accountType
}

var AccountTypeCodeToAccountTypeMap = map[AccountTypeCode]AccountType{
	AccountTypeCodeSavings:  AccountTypeSavings,
	AccountTypeCodeChecking: AccountTypeChecking,
	AccountTypeCodeCurrency: AccountTypeCurrency,
	AccountTypeCodeSalary:   AccountTypeSalary,
	AccountTypeCodeBusiness: AccountTypeBusiness,
}
