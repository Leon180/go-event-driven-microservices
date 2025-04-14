package customizegrpc

import "google.golang.org/grpc"

type GRPCUnaryInterceptor interface {
	Handle() grpc.UnaryServerInterceptor
}

type GRPCStreamInterceptor interface {
	Handle() grpc.StreamServerInterceptor
}
