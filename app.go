package main

import (
	"Protey_microservice/config"
	log_pkg "Protey_microservice/pkg/logger"
	"github.com/spf13/viper"
)

func main() {

	logger := log_pkg.InitLogger()

	logger.Info.Print("Executing InitConfig.")

	if err := config.InitConfig(); err != nil {
		logger.Err.Fatalf(err.Error())
	}

	logger.Info.Print("PORT: ", viper.GetString("PORT"))
	logger.Info.Print("POSTGRES PORT: ", viper.GetString("POSTGRES_PORT"))
	logger.Info.Print("POSTGRES USER: ", viper.GetString("POSTGRES_USER"))

}
