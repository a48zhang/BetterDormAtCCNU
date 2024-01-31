package conf

import (
	"github.com/spf13/viper"
	"main/pkg/logger"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logger.FatalLog(err.Error())
	}
	viper.AutomaticEnv()
}

func GetConf(key string) string {
	return viper.GetString(key)
}
