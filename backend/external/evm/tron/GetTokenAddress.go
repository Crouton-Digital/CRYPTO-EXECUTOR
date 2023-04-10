package tron

import "github.com/sirupsen/logrus"

func GetTokenAddress(token string) string {
	switch token {
	case "USDT":
		return "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	}
	logrus.Errorf("Unknow Token name: %s", token)
	return "nil"
}
