package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.WithError(err).Fatal("config file error")
	}
	viper.WatchConfig()

	log.SetFormatter(&log.TextFormatter{
		TimestampFormat:        time.RFC3339Nano,
		FullTimestamp:          true,
		DisableLevelTruncation: true,
	})
}
