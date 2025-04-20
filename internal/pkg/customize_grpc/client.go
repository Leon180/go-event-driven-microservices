package customizegrpc

import googleGrpc "google.golang.org/grpc"

type GRPCClient interface {
	GetConnection() *googleGrpc.ClientConn
	Close() error
}
