package tron

import (
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"math/big"
)

func GetTronBalance(address string, trc20Contract string) (*big.Int, error) {

	conn := client.NewGrpcClient("grpc.trongrid.io:50051")
	err := conn.Start(grpc.WithInsecure())
	if err != nil {
		logrus.Errorf("Error connecting GRPC Client: %v", err)
	}

	balance, err := conn.TRC20ContractBalance(address, trc20Contract)
	logrus.Infoln(balance)

	return balance, err
}
