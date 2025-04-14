package customizegrpc

type GRPCServer interface {
	RegistGRPCService(registers ...GRPCServiceRegister)
	AddUnaryInterceptors(interceptors ...GRPCUnaryInterceptor)

	Run() error
	GracefulShutdown()
	GetConfig() GRPCConfig
}
