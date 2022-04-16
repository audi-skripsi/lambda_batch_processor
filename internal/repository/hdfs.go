package repository

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func (r *repository) CreateHDFSDirectory(path string) (err error) {
	err = r.hdfsClient.MkdirAll(path, 0777)
	if err != nil {
		r.logger.Errorf("error making directory of %s: %+v", path, err)
	}

	return
}

func (r *repository) WriteToHDFSFile(fileName string, content []byte) (err error) {
	w, err := r.hdfsClient.CreateFile(fileName, 1, 1048576*64, 0777)
	if err != nil {
		r.logger.Errorf("error getting writer to write file: %+v", err)
		return
	}
	defer w.Close()

	_, err = w.Write(content)
	_, err = w.Write([]byte("\n"))
	if err != nil {
		r.logger.Errorf("error writing to file: %+v", err)
	}
	return
}

func (r *repository) AppendToHDFSFile(fileName string, content []byte) (err error) {
	w, err := r.hdfsClient.Append(fileName)
	if err != nil {
		err = r.WriteToHDFSFile(fileName, content)
		if err != nil {
			r.logger.Errorf("error getting writer to append file: %+v", err)
		}
		return
	}
	defer w.Close()

	_, err = w.Write(content)
	_, err = w.Write([]byte("\n"))
	if err != nil {
		r.logger.Errorf("error writing to file: %+v", err)
	}

	return
}

func (r *repository) GetFilePathsWithTimeRange(dir string, startTime time.Time, endTime time.Time) (filePaths []string, err error) {
	filesInfo, err := r.hdfsClient.ReadDir(dir)
	if err != nil {
		r.logger.Errorf("error reading directory: %+v", err)
		return
	}

	for _, v := range filesInfo {
		fileTimeStr := strings.Split(v.Name(), ".txt")
		if len(fileTimeStr) != 2 {
			continue
		}
		fileTimeInt64, err := strconv.ParseInt(fileTimeStr[0], 10, 64)
		if err != nil {
			return nil, err
		}

		fileTime := time.Unix(fileTimeInt64, 0).UTC()
		if fileTime.Before(startTime) || fileTime.After(endTime) {
			continue
		}

		filePaths = append(filePaths, fmt.Sprintf("%s/%s", dir, v.Name()))
	}

	return
}

func (r repository) ReadFile(filePath string) (file []byte, err error) {
	file, err = r.hdfsClient.ReadFile(filePath)
	if err != nil {
		r.logger.Errorf("error reading file: %+v", err)
	}

	return
}
