package config

type KafkaConfig struct {
	Address       string
	ConsumerGroup string
	InTopic       string
	OutTopic      string
}
