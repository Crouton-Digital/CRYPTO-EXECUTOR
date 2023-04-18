package server

import (
	"crypto-executor/internal/cli_menu"
	"crypto-executor/internal/server/config"
	"github.com/sirupsen/logrus"
)

func RunServer() {
	config.LoadServerConfig()

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	logrus.SetReportCaller(config.Config.ServerConfig.LogReportCaller)

	logrus.Info("TRY LOAD CONFIG APP")
	switch config.Config.ServerConfig.DebugLevel {
	case "INFO":
		logrus.SetLevel(logrus.InfoLevel)
	case "ERROR":
		logrus.SetLevel(logrus.ErrorLevel)
	case "DEBUG":
		logrus.SetLevel(logrus.DebugLevel)
	case "FATAL":
		logrus.SetLevel(logrus.FatalLevel)
	case "PANIC":
		logrus.SetLevel(logrus.PanicLevel)
	case "TRACE":
		logrus.SetLevel(logrus.TraceLevel)
	}

	logrus.Info("INFO")
	logrus.Error("ERROR")
	logrus.Debug("DEBUG")
	logrus.Trace("TRACE")
	//logrus.Fatal("FATAL")
	//logrus.Panic("PANIC")

	logrus.Debug("CONFIG")
	logrus.Debug(config.Config)

	//ethereum.EthGenerateWalletWithMnemonic()

	menu := cli_menu.Menu{
		Promt:     "Test Menu",
		CursorPos: 0,
	}

	menu.AddItem("Red", "red")
	menu.AddItem("Green", "green")

	err := menu.Start()
	if err != nil {
		logrus.Errorf("Crash menu: ", err)
	}
	//evm.CreateAccount("testTronAddres")
	//evm.EthGenerateWallet()
	//tron.GetAccountDetailedInfo("TEDbDjEoVeX2qSBvLdbqGFSYG9SbMcELBy")
	//logrus.Info(tron.GetTronBalance("TEDbDjEoVeX2qSBvLdbqGFSYG9SbMcELBy", tron.GetTokenAddress("USDT")))

	//tronClient := tronGRPC.NewGrpcClient("")
	//err := tronClient.Connect(grpc.WithInsecure())
	//if err != nil {
	//	logrus.Errorf("Connecting GRPC Client: %v", err)
	//}

	//WalletClient := tronClient.Client.GetAccount()
	//tronio.
	//go nodemonitoring.Run()
	//StartRouter()
}
