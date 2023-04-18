package tronio

import (
	api "crypto-executor/external/evm/tronio/api"
	"google.golang.org/grpc"
	"time"
)

type GrpcClient struct {
	Address     string
	Conn        *grpc.ClientConn
	Client      api.WalletClient
	grpcTimeout time.Duration
	opts        []grpc.DialOption
	apiKey      string
}

func NewGrpcClient(address string) *GrpcClient {
	client := &GrpcClient{
		Address:     address,
		grpcTimeout: 5 * time.Second,
	}
	return client
}

func (g *GrpcClient) SetTimeout(timeout time.Duration) {
	g.grpcTimeout = timeout
}

func (g *GrpcClient) SetAPIKey(apiKey string) {
	g.apiKey = apiKey
}

func (g *GrpcClient) Connect(opts ...grpc.DialOption) error {
	var err error
	if len(g.Address) < 0 {
		g.Address = "grpc.trongrid.io:50051"
	}

	g.opts = opts

	g.Conn, err = grpc.Dial(g.Address, opts...)
	if err != nil {
		return err
	}

	g.Client = api.NewWalletClient(g.Conn)

	return nil
}
