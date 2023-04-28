package config

import (
	"github.com/sirupsen/logrus"
	yaml3 "gopkg.in/yaml.v3"
	"os"
)

type AppConfig struct {
	ServerConfig Server `yaml:"server"`
	OKX          OKX    `yaml:"okx"`
}

type Server struct {
	HttpPort        string `yaml:"http_port"`
	DebugLevel      string `yaml:"debug"`
	LogReportCaller bool   `yaml:"log_report_caller"`
}

type OKX struct {
	ApiKey     string `yaml:"api_key"`
	SecretKey  string `yaml:"secret_key"`
	PassPhrase string `yaml:"passphrase"`
}

var (
	Config AppConfig
)

func LoadServerConfig() {

	env := os.Getenv("ENV")

	if env == "" {
		env = "local"
	}

	data, err := os.ReadFile("config/" + env + ".yml")
	if err != nil {
		logrus.Errorf("Failed to read config: %v", err)
		os.Exit(1)
	}
	err = yaml3.Unmarshal(data, &Config)
	if err != nil {
		logrus.Errorf("Failed to parse config: %v", err)
		os.Exit(1)
	}

	////logrus.Info(Config)
	//for key, network_nodes := range Config.NetworksConfig {
	//	logrus.Infof("======== %v ========", key)
	//	for _, network_node := range network_nodes.Nodes {
	//		logrus.Infof("%v | %v %v Public: %v", key, network_node.Name, network_node.Url, network_node.Public)
	//	}
	//}
}
