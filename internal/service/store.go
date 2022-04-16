package service

import (
	"github.com/audi-skripsi/lambda_batch_processor/internal/util/converterutil"
	"github.com/audi-skripsi/lambda_batch_processor/pkg/dto"
)

func (s *service) StoreToDataLake(event dto.EventLog) (err error) {
	b64, err := converterutil.EventLogDtoToBase64(event)
	if err != nil {
		s.logger.Errorf("error decoding to base64 of %s: %+v", event.UID, err)
		return
	}
	s.logger.Print(b64)
	return
}
