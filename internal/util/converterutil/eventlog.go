package converterutil

import (
	"encoding/base64"
	"encoding/json"

	"github.com/audi-skripsi/lambda_batch_processor/internal/model"
	"github.com/audi-skripsi/lambda_batch_processor/pkg/dto"
)

func EventLogDtoToModel(logDto dto.EventLog) (logModel model.EventLog) {
	return model.EventLog{
		UID:       logDto.UID,
		Level:     logDto.Level,
		AppName:   logDto.AppName,
		Timestamp: logDto.Timestamp,
		Data:      logDto.Data,
	}
}

func EventLogDtoToBase64(logDto dto.EventLog) (b64 string, err error) {
	b, err := json.Marshal(logDto)
	if err != nil {
		return
	}

	b64 = base64.StdEncoding.EncodeToString(b)
	return
}

func Base64ToEventLogDto(b64 string) (logDto dto.EventLog, err error) {
	var dst []byte

	_, err = base64.StdEncoding.Decode(dst, []byte(b64))
	if err != nil {
		return
	}

	err = json.Unmarshal(dst, &logDto)
	return
}
