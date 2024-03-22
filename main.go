package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type Configuration struct {
	App     AppConfiguration
	MysqlDb MysqlConfiguration
	MongoDb MongoDBConfiguration
	Redis   RedisCofiguration
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

type RedisCofiguration struct {
	Address  string
	Password string
	Db       int
}

func main() {
	ctx := context.Background()
	// === Set up Config ===
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var config Configuration
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	// === Mysql DB ===

	// === Redis ===
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
		DB:       config.Redis.Db,
	})

	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("Error connect Redis, %s", err)
	}

}
