package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName     string
	AppAddress  string
	KafkaConfig KafkaConfig
	HDFSConfig  HDFSConfig
}

var config *Config

func Init() {
	err := godotenv.Load("conf/.env")
	if err != nil {
		log.Printf("[Init] error on loading env from file: %+v", err)
	}

	config = &Config{
		AppName:    os.Getenv("APP_NAME"),
		AppAddress: os.Getenv("APP_ADDRESS"),
		KafkaConfig: KafkaConfig{
			Address:       os.Getenv("KAFKA_ADDRESS"),
			InTopic:       os.Getenv("KAFKA_IN_TOPIC"),
			OutTopic:      os.Getenv("KAFKA_OUT_TOPIC"),
			ConsumerGroup: os.Getenv("KAFKA_CONSUMER_GROUP"),
		},
		HDFSConfig: HDFSConfig{
			NameNodeAddress: os.Getenv("NAME_NODE_ADDRESS"),
		},
	}

	if config.AppName == "" {
		log.Panicf("[Init] app name cannot be empty")
	}

	if config.AppAddress == "" {
		log.Panicf("[Init] app address cannot be empty")
	}

	if config.KafkaConfig.Address == "" ||
		config.KafkaConfig.InTopic == "" ||
		config.KafkaConfig.OutTopic == "" {
		log.Panicf("[Init] kafka config cannot be empty")
	}

	if config.HDFSConfig.NameNodeAddress == "" {
		log.Panicf("[Init] namenode address cannot be empty")
	}
}

func Get() *Config {
	return config
}
