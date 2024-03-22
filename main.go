package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	App     AppConfiguration
	MysqlDb MysqlConfiguration
	MongoDb MongoDBConfiguration
}
type AppConfiguration struct {
	Env   string
	Port  uint
	Debug bool
}
type MysqlConfiguration struct {
	Host     string
	Port     uint
	User     string
	Password string
	Database string
}

type MongoDBConfiguration struct {
	Connection string
}

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var configuration Configuration
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	fmt.Println(configuration.App)
	fmt.Println(configuration.MysqlDb)
	fmt.Println(configuration.MongoDb)

}
