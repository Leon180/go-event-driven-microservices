package enumsaccounts

type AccountType string

const (
	AccountTypeInvalid AccountType = "invalid"
	AccountTypeSavings AccountType = "saving"
	AccountTypeCurrent AccountType = "current"
	AccountTypeSalary  AccountType = "salary"
)

func (a AccountType) IsValid() bool {
	active, ok := AccountTypActiveMap[a]
	return ok && active
}

var AccountTypActiveMap = map[AccountType]bool{
	AccountTypeSavings: true,
	AccountTypeCurrent: true,
	AccountTypeSalary:  true,
}

func (a AccountType) ToAccountTypeCode() AccountTypeCode {
	code, ok := AccountTypeToCodeMap[a]
	if !ok {
		return AccountTypeCodeInvalid
	}
	return code
}

var AccountTypeToCodeMap = map[AccountType]AccountTypeCode{
	AccountTypeSavings: AccountTypeCodeSavings,
	AccountTypeCurrent: AccountTypeCodeCurrent,
	AccountTypeSalary:  AccountTypeCodeSalary,
}

type AccountTypeCode int

const (
	AccountTypeCodeInvalid AccountTypeCode = iota - 1
	AccountTypeCodeSavings
	AccountTypeCodeCurrent
	AccountTypeCodeSalary
)

func (a AccountTypeCode) ToAccountType() AccountType {
	accountType, ok := AccountTypeCodeToAccountTypeMap[a]
	if !ok {
		return AccountTypeInvalid
	}
	return accountType
}

var AccountTypeCodeToAccountTypeMap = map[AccountTypeCode]AccountType{
	AccountTypeCodeSavings: AccountTypeSavings,
	AccountTypeCodeCurrent: AccountTypeCurrent,
	AccountTypeCodeSalary:  AccountTypeSalary,
}
