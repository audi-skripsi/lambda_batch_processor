package service

import (
	"github.com/audi-skripsi/lambda_batch_processor/internal/config"
	"github.com/audi-skripsi/lambda_batch_processor/internal/repository"
	"github.com/audi-skripsi/lambda_batch_processor/pkg/dto"
	"github.com/sirupsen/logrus"
)

type Service interface {
	Ping() (pingResponse string, timestamp int64)
	StoreToDataLake(event dto.EventLog) (err error)
	ExtractEvents(req dto.EventExtractionRequest) (total int, err error)
}

type service struct {
	logger     *logrus.Entry
	repository repository.Repository
	config     *serviceConfig
}

type serviceConfig struct {
	KafkaConfig *config.KafkaConfig
}

type NewServiceParams struct {
	Logger     *logrus.Entry
	Repository repository.Repository
	Config     *config.Config
}

func NewService(params NewServiceParams) Service {
	return &service{
		logger:     params.Logger,
		repository: params.Repository,
		config: &serviceConfig{
			KafkaConfig: &params.Config.KafkaConfig,
		},
	}
}
