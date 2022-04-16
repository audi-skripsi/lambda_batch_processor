package component

import (
	"github.com/audi-skripsi/lambda_batch_processor/internal/config"
	"github.com/colinmarc/hdfs/v2"
)

func NewHDFSClient(config config.HDFSConfig) (hdfsClient *hdfs.Client, err error) {
	hdfsClient, err = hdfs.New(config.NameNodeAddress)
	return
}
