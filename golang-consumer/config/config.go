package config

import (
	"golang-consumer/models"
	"log"

	"github.com/spf13/viper"
)

func New() models.Config {
	v := viper.New()
	v.SetConfigFile("./config/config.json")
	v.SetConfigType("json")

	err := v.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}

	var config models.Config
	if e := v.Unmarshal(&config); e != nil {
		log.Fatal(err)
	}

	return config

}
