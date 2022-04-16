package handler

import (
	"net/http"

	"github.com/audi-skripsi/lambda_batch_processor/pkg/dto"
	"github.com/audi-skripsi/lambda_batch_processor/pkg/util/httputil"
)

type PingService func() (string, int64)

func HandlePing(service PingService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg, timestamp := service()
		httputil.WriteSuccessResponse(w, dto.PublicPingResponse{
			Message:   msg,
			Timestamp: timestamp,
		})
	}
}
