package api

import (
	context "context"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"google.golang.org/grpc"
)

const _ = grpc.SupportPackageIsVersion7

type WalletClient interface {
	GetAccount(ctx context.Context, in *core.Account, opts ...grpc.CallOption) (*core.Account, error)
}

type walletClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletClient(cc grpc.ClientConnInterface) WalletClient {
	return &walletClient{cc}
}

func (c *walletClient) GetAccount(ctx context.Context, in *core.Account, opts ...grpc.CallOption) (*core.Account, error) {
	out := new(core.Account)
	err := c.cc.Invoke(ctx, "/protocol.Wallet/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
