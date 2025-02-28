package enums

import (
	"regexp"
	"strconv"
	"time"
)

type Format string

const (
	MobileNumberFormat Format = "^[0-9]{10}$"
)

func (f Format) String() string {
	return string(f)
}

func (f Format) ValidateFormatString(st string) bool {
	matched, _ := regexp.MatchString(f.String(), st)
	return matched
}

func (f Format) ValidateFormat(st any) bool {
	switch v := st.(type) {
	case string:
		return f.ValidateFormatString(v)
	case int:
		return f.ValidateFormatString(strconv.Itoa(v))
	case int64:
		return f.ValidateFormatString(strconv.FormatInt(v, 10))
	case float64:
		return f.ValidateFormatString(strconv.FormatFloat(v, 'f', -1, 64))
	case bool:
		return f.ValidateFormatString(strconv.FormatBool(v))
	case []byte:
		return f.ValidateFormatString(string(v))
	case time.Time:
		return f.ValidateFormatString(v.Format(time.RFC3339))
	default:
		return false
	}
}
