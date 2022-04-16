package service

import (
	"bufio"
	"bytes"

	"github.com/audi-skripsi/lambda_batch_processor/internal/constant"
	"github.com/audi-skripsi/lambda_batch_processor/internal/util/converterutil"
	"github.com/audi-skripsi/lambda_batch_processor/internal/util/timeutil"
	"github.com/audi-skripsi/lambda_batch_processor/pkg/dto"
	"github.com/audi-skripsi/lambda_batch_processor/pkg/errors"
)

func (s *service) ExtractEvents(req dto.EventExtractionRequest) (total int, err error) {
	startTime, err := timeutil.TimeStringToTime(req.StartTime)
	if err != nil {
		s.logger.Errorf("error parsing time: %+v", err)
		err = errors.ErrBadRequest
		return
	}

	endTime, err := timeutil.TimeStringToTime(req.EndTime)
	if err != nil {
		s.logger.Errorf("error parsing time: %+v", err)
		err = errors.ErrBadRequest
		return
	}

	filePaths, err := s.repository.GetFilePathsWithTimeRange(constant.HDFSDataLakeBasePath, startTime, endTime)
	if err != nil {
		s.logger.Errorf("error get file paths from dir: %+v", err)
		err = errors.ErrInternalServer
		return
	}

	for _, filePath := range filePaths {
		file, err := s.repository.ReadFile(filePath)
		if err != nil {
			s.logger.Errorf("error reading from file: %+v", err)
			break
		}

		line := bufio.NewScanner(bytes.NewReader(file))
		for line.Scan() {
			if line.Err() != nil {
				s.logger.Errorf("error reading file: %+v", err)
				break
			}
			var eventLog dto.EventLog

			eventLog, err = converterutil.Base64ToEventLogDto(line.Text())
			if err != nil {
				s.logger.Errorf("error converting to event log: %+v", err)
				break
			}

			eventLogModel := converterutil.EventLogDtoToModel(eventLog)
			err = s.repository.PushEventLog(eventLogModel)
			if err != nil {
				s.logger.Errorf("error pushing to kafka: %+v", err)
				break
			}
			total++
		}
	}

	return
}
