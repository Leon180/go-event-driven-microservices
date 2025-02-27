package enums

type ServiceName string

const (
	ServiceNameAccount ServiceName = "account"
)

func (s ServiceName) ToString() string {
	return string(s)
}
