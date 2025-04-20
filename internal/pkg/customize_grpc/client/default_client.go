package customizegrpcclient

import (
	"net"

	customizegrpc "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_grpc"
	googleGrpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCClient(
	config customizegrpc.GRPCConfig,
) (customizegrpc.GRPCClient, error) {
	conn, err := googleGrpc.NewClient(
		net.JoinHostPort(config.GetHost(), config.GetPort()),
		googleGrpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	return &grpcClientImpl{
		config: config,
		conn:   conn,
	}, nil
}

type grpcClientImpl struct {
	conn   *googleGrpc.ClientConn
	config customizegrpc.GRPCConfig
}

func (c *grpcClientImpl) GetConnection() *googleGrpc.ClientConn {
	return c.conn
}

func (g *grpcClientImpl) Close() error {
	return g.conn.Close()
}
