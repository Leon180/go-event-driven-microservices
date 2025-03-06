package enums

type TimeFormat string

const (
	TimeFormatClockOnly TimeFormat = "15:04"
)

func (t TimeFormat) ToString() string {
	return string(t)
}
