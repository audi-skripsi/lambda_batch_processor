package handler

import (
	"net/http"

	"github.com/audi-skripsi/lambda_batch_processor/pkg/dto"
	"github.com/audi-skripsi/lambda_batch_processor/pkg/util/httputil"
	"github.com/audi-skripsi/lambda_batch_processor/pkg/util/jsonutil"
)

type ExtractorFunc func(param dto.EventExtractionRequest) (total int, err error)

func HandleExtraction(handlefunc ExtractorFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.EventExtractionRequest

		err := jsonutil.ConvertToObject(r, &req)
		if err != nil {
			httputil.WriteErrorResponse(w, err)
			return
		}

		total, err := handlefunc(req)
		if err != nil {
			httputil.WriteErrorResponse(w, err)
			return
		}

		httputil.WriteSuccessResponse(w, dto.EventExtractionResponse{TotalExtracted: total})
	}
}
