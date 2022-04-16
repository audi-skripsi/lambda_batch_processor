package repository

func (r *repository) CreateHDFSDirectory(path string) (err error) {
	err = r.hdfsClient.MkdirAll(path, 0777)
	if err != nil {
		r.logger.Errorf("error making directory of %s: %+v", path, err)
	}

	return
}
