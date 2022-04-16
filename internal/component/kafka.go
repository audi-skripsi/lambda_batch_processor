package component

import (
	"github.com/audi-skripsi/lambda_batch_processor/internal/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaPublisher(config config.KafkaConfig) (producer *kafka.Producer, err error) {
	producer, err = kafka.NewProducer(
		&kafka.ConfigMap{"bootstrap.servers": config.Address},
	)
	return
}

func NewKafkaConsumer(config config.KafkaConfig) (consumer *kafka.Consumer, err error) {
	conf := kafka.ConfigMap{}
	conf["bootstrap.servers"] = config.Address
	if config.ConsumerGroup != "" {
		conf["group.id"] = config.ConsumerGroup
	}
	conf["auto.offset.reset"] = "smallest"
	conf["enable.auto.commit"] = false

	consumer, err = kafka.NewConsumer(&conf)
	return
}
