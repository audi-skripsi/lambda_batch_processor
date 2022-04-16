package http

import (
	"net/http"

	"github.com/audi-skripsi/lambda_batch_processor/cmd/http/router"
	"github.com/audi-skripsi/lambda_batch_processor/internal/config"
	"github.com/audi-skripsi/lambda_batch_processor/internal/service"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type ServerInitParams struct {
	Logger  *logrus.Entry
	Config  *config.Config
	Service service.Service
}

func StartServer(param ServerInitParams) {
	param.Logger.Info("starting service in http mode...")

	ro := mux.NewRouter()

	router.Init(router.RouterInitParams{
		Router:  ro,
		Service: param.Service,
	})

	go func() {
		err := http.ListenAndServe(param.Config.AppAddress, ro)
		if err != nil {
			param.Logger.Errorf("error listening to http: %v", err)
		}
	}()
}
