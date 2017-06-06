package utils

import (
	"github.com/spf13/viper"
)

const ModeProduction = "production"
const ModeDevelopment = "development"

type Configuration struct {
		ServiceURL string
    ServicePort string
		ServiceMode string
    DatabaseHost  string
		DatabasePort int
		DatabaseDatabase  string
		DatabaseUser  string
		DatabasePassword  string
}

func Load() Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	
	var configuration Configuration

	err := viper.ReadInConfig()
	if err != nil {
		ErrorLog.Println("Config file not found..")
	} else {
		configuration.ServiceURL = viper.GetString("service.url")
		configuration.ServicePort = viper.GetString("service.port")
		configuration.ServiceMode = viper.GetString("service.mode")

		configuration.DatabaseHost = viper.GetString("database.host")
		configuration.DatabasePort = viper.GetInt("database.port")
		configuration.DatabaseDatabase = viper.GetString("database.database")
		configuration.DatabaseUser = viper.GetString("database.user")
		configuration.DatabasePassword = viper.GetString("database.password")

		ErrorLog.Printf("\nUsing configuration:\n service port = %s\n database host = %s\n", 
			configuration.ServicePort, configuration.DatabaseHost)
	}
	return configuration
}
