package tron

import (
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/account"
	"github.com/fbsobreira/gotron-sdk/pkg/store"
	"github.com/sirupsen/logrus"
)

func TronCreateAccount() {

	acc := &account.Creation{
		Name:       "maxiWallet",
		Passphrase: "opdjkshdfshfuwefijui",
	}

	account.CreateNewLocalAccount(acc)
	logrus.Info(acc)
	addr, _ := store.AddressFromAccountName(acc.Name)
	fmt.Printf("Tron Address: %s\n", addr)

	store.DescribeLocalAccounts()

}
