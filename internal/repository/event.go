package repository

import (
	"encoding/json"

	"github.com/audi-skripsi/lambda_batch_processor/internal/model"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (r *repository) PushEventLog(log model.EventLog) (err error) {
	var msg kafka.Message
	msg.Key = []byte(log.AppName)

	b, _ := json.Marshal(log)

	msg.Value = b

	r.logger.Infof("logging to %s with message of %+v", r.config.kafkaConfig.OutTopic, log)

	msg.TopicPartition = kafka.TopicPartition{
		Topic:     &r.config.kafkaConfig.OutTopic,
		Partition: kafka.PartitionAny,
	}

	err = r.kafkaProducer.Produce(&msg, nil)
	if err != nil {
		r.logger.Errorf("error on pushing to kafka: %+v", err)
	}
	return
}
