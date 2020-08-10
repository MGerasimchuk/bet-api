package config

import (
	"os"

	"github.com/spf13/viper"
)

type (
	LogConfig struct {
		Level int
	}
	DBConfig struct {
		Host     string
		Port     int
		Name     string
		User     string
		Password string
	}
	HTTPServerConfig struct {
		Port int
		Mode string
	}
	Canceller struct {
		CancelTransactionsCount        int
		RepeatCancellationEveryMinutes float64
	}

	RootConfig struct {
		Log        LogConfig
		DB         DBConfig
		HTTPServer HTTPServerConfig
		Canceller  Canceller
	}
)

func GetRootConfig() RootConfig {
	viper.SetConfigFile(os.Getenv("CONFIG_FILE"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	cfg := RootConfig{}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
