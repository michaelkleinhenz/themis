package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
    ServicePort int
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
		fmt.Println("Config file not found..")
	} else {
		configuration.ServicePort = viper.GetInt("service.port")

		configuration.DatabaseHost = viper.GetString("database.host")
		configuration.DatabasePort = viper.GetInt("database.port")
		configuration.DatabaseDatabase = viper.GetString("database.database")
		configuration.DatabaseUser = viper.GetString("database.user")
		configuration.DatabasePassword = viper.GetString("database.password")

		fmt.Printf("\nUsing configuration:\n service port = %d\n database host = %s\n", 
			configuration.ServicePort, configuration.DatabaseHost)
	}
	return configuration
}
