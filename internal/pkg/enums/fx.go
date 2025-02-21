package enums

type FxGroup string

const (
	FxGroupMiddlewares FxGroup = "middlewares"
	FxGroupEndpoints   FxGroup = "endpoints"
)

func (g FxGroup) ToString() string {
	return string(g)
}
