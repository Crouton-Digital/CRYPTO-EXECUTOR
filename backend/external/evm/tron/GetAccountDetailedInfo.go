package tron

import (
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func GetAccountDetailedInfo(address string) {

	conn := client.NewGrpcClient("grpc.trongrid.io:50051")
	err := conn.Start(grpc.WithInsecure())
	if err != nil {
		logrus.Errorf("Error connecting GRPC Client: %v", err)
	}
	acc, err := conn.GetAccountDetailed(address)
	logrus.Info(acc.Assets)
	logrus.Infoln(acc)

}
