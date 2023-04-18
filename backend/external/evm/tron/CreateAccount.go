package tron

import (
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/account"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/store"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func CreateAccount(WalletName string) {

	conn := client.NewGrpcClient("grpc.trongrid.io:50051")
	err := conn.Start(grpc.WithInsecure())
	if err != nil {
		logrus.Errorf("Error connecting GRPC Client: %v", err)
	}

	acc := &account.Creation{
		Name:       WalletName,
		Passphrase: "opdjkshdfshfuwefijui",
	}

	account.CreateNewLocalAccount(acc)
	logrus.Info(acc)
	addr, _ := store.AddressFromAccountName(acc.Name)
	fmt.Printf("Tron Address: %s\n", addr)

	store.DescribeLocalAccounts()
	account.RemoveAccount(WalletName)

}
