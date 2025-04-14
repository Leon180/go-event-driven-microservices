package enums

type FxGroup string

const (
	FxGroupMiddlewares           FxGroup = "middlewares"
	FxGroupEndpoints             FxGroup = "endpoints"
	FxGroupGRPCUnaryInterceptors FxGroup = "grpcserverinterceptors"
	FxGroupGRPCServiceRegister   FxGroup = "grpcserverregister"
)

func (g FxGroup) ToString() string {
	return string(g)
}
