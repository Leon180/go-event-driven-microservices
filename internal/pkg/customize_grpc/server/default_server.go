package customizegrpcserver

import (
	"net"

	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"

	customizegrpc "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_grpc"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcmiddlewarerecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcmiddlewaretags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	googleGrpc "google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func NewGRPCServer(
	config customizegrpc.GRPCConfig,
	logger loggers.Logger,
) customizegrpc.GRPCServer {
	return &grpcServerImpl{
		config: config,
		logger: logger,
		unaryInterceptors: []googleGrpc.UnaryServerInterceptor{
			grpcmiddlewaretags.UnaryServerInterceptor(),
			grpcmiddlewarerecovery.UnaryServerInterceptor(),
		},
	}
}

type grpcServerImpl struct {
	server            *googleGrpc.Server
	config            customizegrpc.GRPCConfig
	logger            loggers.Logger
	unaryInterceptors []googleGrpc.UnaryServerInterceptor
	grpcServices      []customizegrpc.GRPCServiceRegister
}

func (s *grpcServerImpl) AddUnaryInterceptors(interceptors ...customizegrpc.GRPCUnaryInterceptor) {
	for _, interceptor := range interceptors {
		s.unaryInterceptors = append(s.unaryInterceptors, interceptor.Handle())
	}
}

// RegistGRPCService registers the GRPC service for the GRPC server
// for example:
//
//	register := func(s *googleGrpc.Server) {
//		helloworld.RegisterGreeterServer(s, grpcService)
//	}
//	s.RegistGRPCService(register)
func (s *grpcServerImpl) RegistGRPCService(registers ...customizegrpc.GRPCServiceRegister) {
	s.grpcServices = append(s.grpcServices, registers...)
}

func (s *grpcServerImpl) Run() error {
	l, err := net.Listen("tcp", net.JoinHostPort(s.config.GetHost(), s.config.GetPort()))
	if err != nil {
		return err
	}

	s.server = googleGrpc.NewServer(
		googleGrpc.KeepaliveParams(s.config.GetKeepAliveParams()),
		googleGrpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(s.unaryInterceptors...)),
	)
	for _, register := range s.grpcServices {
		register(s.server)
	}
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s.server, healthServer)
	healthServer.SetServingStatus(
		s.config.GetServiceName(),
		grpc_health_v1.HealthCheckResponse_SERVING,
	)

	if s.config.GetEnvironment() == enums.EnvironmentDevelopment {
		reflection.Register(s.server)
	}

	if err := s.server.Serve(l); err != nil {
		return err
	}

	return nil
}

func (s *grpcServerImpl) GracefulShutdown() {
	s.server.Stop()
	s.server.GracefulStop()
}

func (s *grpcServerImpl) GetConfig() customizegrpc.GRPCConfig {
	return s.config
}
