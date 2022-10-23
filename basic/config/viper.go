package config

import (
	"github.com/spf13/viper"
	setup "github.com/tianshengdiwangK/simProject/basic/log"
	"go.uber.org/zap"
	"sync"
)

var (
	once   sync.Once
	config *viper.Viper
	logger *zap.SugaredLogger
)

func init() {
	once.Do(func() {
		logger = setup.CwLog()
		config = viper.New()
		config.AddConfigPath("././")
		config.SetConfigName("config")
		config.SetConfigType("yaml")
		if err := config.ReadInConfig(); err == nil {
			logger.Infof("Read config successfully: %s", config.ConfigFileUsed())
		} else {
			logger.Fatalf("Read failed: %s \n", err)
		}
	},
	)
}
func NewViper() *viper.Viper {
	return config
}
