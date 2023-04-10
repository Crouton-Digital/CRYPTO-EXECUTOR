package server

import (
	"crypto-executor/external/evm/tron"
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

	//logrus.Info("INFO")
	//logrus.Error("ERROR")
	//logrus.Debug("DEBUG")
	//logrus.Fatal("FATAL")
	//logrus.Panic("PANIC")
	//logrus.Trace("TRACE")

	logrus.Debug("CONFIG")
	logrus.Debug(config.Config)

	//evm.TronCreateAccount()
	//evm.EthGenerateWallet()
	tron.GetAccountDetailedInfo("TEDbDjEoVeX2qSBvLdbqGFSYG9SbMcELBy")
	//logrus.Info(tron.GetTronBalance("TEDbDjEoVeX2qSBvLdbqGFSYG9SbMcELBy", tron.GetTokenAddress("USDT")))
	//go nodemonitoring.Run()
	//StartRouter()
}
