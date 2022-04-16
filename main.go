package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/audi-skripsi/lambda_batch_processor/cmd/consumer"
	"github.com/audi-skripsi/lambda_batch_processor/internal/component"
	"github.com/audi-skripsi/lambda_batch_processor/internal/config"
	"github.com/audi-skripsi/lambda_batch_processor/internal/repository"
	"github.com/audi-skripsi/lambda_batch_processor/internal/service"
	"github.com/audi-skripsi/lambda_batch_processor/pkg/util/logutil"
)

func main() {
	config.Init()
	config := config.Get()

	logger := logutil.NewLogger(logutil.NewLoggerParams{
		PrettyPrint: true,
		ServiceName: config.AppName,
	})

	logger.Infof("app initialized with config of: %+v", config)

	kafkaConsumer, err := component.NewKafkaConsumer(config.KafkaConfig)
	if err != nil {
		logger.Fatalf("[main] error initializing kafka consumer: %+v", err)
	}

	kafkaProducer, err := component.NewKafkaPublisher(config.KafkaConfig)
	if err != nil {
		logger.Fatalf("[main] error initializing kafka publisher: %+v", err)
	}

	repository := repository.NewRepository(repository.NewRepositoryParams{
		Logger:        logger,
		KafkaProducer: kafkaProducer,
		Config:        config,
	})

	service := service.NewService(service.NewServiceParams{
		Logger:     logger,
		Repository: repository,
		Config:     config,
	})

	consumer := consumer.NewConsumer(consumer.NewConsumerParams{
		Logger:        logger,
		KafkaConsumer: kafkaConsumer,
		Config:        config,
		Service:       service,
	})

	consumer.Init()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	go func() {
		go kafkaConsumer.Close()
		go kafkaProducer.Close()
	}()
	logger.Info("stopping service gracefully...")
	time.Sleep(2 * time.Second)
	logger.Info("service stopped gracefully")
}
