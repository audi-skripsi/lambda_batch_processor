package repository

import (
	"github.com/audi-skripsi/lambda_batch_processor/internal/config"
	"github.com/audi-skripsi/lambda_batch_processor/internal/model"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	PushEventLog(log model.EventLog) (err error)
}

type repository struct {
	logger        *logrus.Entry
	kafkaProducer *kafka.Producer
	config        *repositoryConfig
}

type repositoryConfig struct {
	kafkaConfig config.KafkaConfig
}

type NewRepositoryParams struct {
	Logger        *logrus.Entry
	KafkaProducer *kafka.Producer
	Config        *config.Config
}

func NewRepository(params NewRepositoryParams) Repository {
	return &repository{
		logger:        params.Logger,
		kafkaProducer: params.KafkaProducer,
		config: &repositoryConfig{
			kafkaConfig: params.Config.KafkaConfig,
		},
	}
}
