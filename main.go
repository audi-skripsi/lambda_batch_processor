package main

import (
	"flag"
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

var logInjectorMode *bool
var logExtractorMode *bool

func init() {
	logInjectorMode = flag.Bool("log-injector", false, "used for ignesting logs from kafka to data lake")
	logExtractorMode = flag.Bool("log-extractor", false, "used for extracting logs from data lake to further processing")
}

func main() {
	flag.Parse()

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

	hdfsClient, err := component.NewHDFSClient(config.HDFSConfig)
	if err != nil {
		logger.Fatalf("[main] error initializing hdfs: %+v", err)
	}

	repository := repository.NewRepository(repository.NewRepositoryParams{
		Logger:        logger,
		KafkaProducer: kafkaProducer,
		Config:        config,
		HDFSClient:    hdfsClient,
	})

	service := service.NewService(service.NewServiceParams{
		Logger:     logger,
		Repository: repository,
		Config:     config,
	})

	if *logInjectorMode {
		logger.Infof("starting app as log-injector mode")
		consumer := consumer.NewConsumer(consumer.NewConsumerParams{
			Logger:        logger,
			KafkaConsumer: kafkaConsumer,
			Config:        config,
			Service:       service,
		})
		consumer.Init()
	} else if *logExtractorMode {
		logger.Infof("starting app as log-extractor mode")
	} else {
		logger.Errorf("invalid app mode, should be log-injector or log-extractor")
	}

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
