package repository

import (
	"github.com/audi-skripsi/lambda_batch_processor/internal/config"
	"github.com/audi-skripsi/lambda_batch_processor/internal/model"
	"github.com/colinmarc/hdfs/v2"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	PushEventLog(log model.EventLog) (err error)

	CreateHDFSDirectory(path string) (err error)
}

type repository struct {
	logger        *logrus.Entry
	kafkaProducer *kafka.Producer
	hdfsClient    *hdfs.Client
	config        *repositoryConfig
}

type repositoryConfig struct {
	kafkaConfig config.KafkaConfig
}

type NewRepositoryParams struct {
	Logger        *logrus.Entry
	KafkaProducer *kafka.Producer
	HDFSClient    *hdfs.Client
	Config        *config.Config
}

func NewRepository(params NewRepositoryParams) Repository {
	return &repository{
		logger:        params.Logger,
		kafkaProducer: params.KafkaProducer,
		hdfsClient:    params.HDFSClient,
		config: &repositoryConfig{
			kafkaConfig: params.Config.KafkaConfig,
		},
	}
}
