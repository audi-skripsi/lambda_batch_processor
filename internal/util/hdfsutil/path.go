package hdfsutil

import (
	"fmt"

	"github.com/audi-skripsi/lambda_batch_processor/internal/constant"
)

func AppendWithBaseDataLakePath(name string) (absPath string) {
	absPath = fmt.Sprintf("%s/%s", constant.HDFSDataLakeBasePath, name)
	return
}
