package main

import (
	"fmt"

	"themis/workitems"

	"gopkg.in/mgo.v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Configuration struct {
    servicePort int
    databaseHost  string
		databasePort int
		databaseDatabase  string
		databaseUser  string
		databasePassword  string
}

func InitConfig() Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	
	var configuration Configuration

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file not found..")
	} else {
		configuration.servicePort = viper.GetInt("service.port")

		configuration.databaseHost = viper.GetString("database.host")
		configuration.databasePort = viper.GetInt("database.port")
		configuration.databaseDatabase = viper.GetString("database.database")
		configuration.databaseUser = viper.GetString("database.user")
		configuration.databasePassword = viper.GetString("database.password")

		fmt.Printf("\nUsing configuration:\n service port = %d\n database host = %s\n", 
			configuration.servicePort, configuration.databaseHost)
	}
	return configuration
}

func ConnectToDatabase(configuration Configuration) *mgo.Database {
	session, err := mgo.DialWithInfo(&mgo.DialInfo {
		Addrs:    []string { configuration.databaseHost },
		Username: configuration.databaseUser,
		Password: configuration.databasePassword,
		Database: configuration.databaseDatabase,
	})
	if err != nil {
		panic(err)
	}
	//defer session.Close()
	fmt.Printf("Connected to database at %v!\n", session.LiveServers())
	return session.DB(configuration.databaseDatabase)
}

func main() {
	configuration := InitConfig()
	database := ConnectToDatabase(configuration)
	workitems.WriteToDatabase(database)
	workItem := workitems.ReadFromDatabase(database, "newId0")
	fmt.Printf("WorkItem data: %s %s", workItem.Id, workItem.Type)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
