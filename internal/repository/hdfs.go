package repository

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
