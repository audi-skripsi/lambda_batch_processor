package httputil

import (
	"encoding/json"
	"net/http"

	"github.com/audi-skripsi/lambda_batch_processor/pkg/dto"
	"github.com/audi-skripsi/lambda_batch_processor/pkg/errors"
)

func WriteSuccessResponse(w http.ResponseWriter, payload interface{}) {
	WriteResponse(w, dto.ResponseParam{
		Status: http.StatusOK,
		Payload: dto.BaseResponse{
			Data: payload,
		},
	})
}

func WriteErrorResponse(w http.ResponseWriter, er error) {
	errResp := errors.GetErrorResponse(er)
	WriteResponse(w, dto.ResponseParam{
		Status: int(errResp.Code),
		Payload: dto.BaseResponse{
			Error: &dto.ErrorResponse{
				Code:    errResp.Code,
				Message: errResp.Message,
			},
		},
	})
}

func WriteResponse(w http.ResponseWriter, param dto.ResponseParam) {
	w.Header().Add("Content-Type", "application/json")
	b, err := json.Marshal(param.Payload)
	if err != nil {
		return
	}
	w.WriteHeader(param.Status)
	w.Write(b)
}
