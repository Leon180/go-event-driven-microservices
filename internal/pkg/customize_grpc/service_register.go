package customizegrpc

import "google.golang.org/grpc"

type GRPCServiceRegister func(s *grpc.Server)
