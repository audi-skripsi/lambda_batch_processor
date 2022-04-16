package router

import (
	"net/http"

	"github.com/audi-skripsi/lambda_batch_processor/cmd/http/handler"
	"github.com/audi-skripsi/lambda_batch_processor/internal/service"
	"github.com/gorilla/mux"
)

type RouterInitParams struct {
	Router  *mux.Router
	Service service.Service
}

func Init(params RouterInitParams) {
	v1Route := params.Router.PathPrefix("/v1").Subrouter()

	v1Route.HandleFunc(
		PingPath,
		handler.HandlePing(params.Service.Ping),
	).Methods(http.MethodGet)

	v1Route.HandleFunc(
		ExtractionPath,
		handler.HandleExtraction(params.Service.ExtractEvents),
	).Methods(http.MethodPost)

}
