package confighup

type StorageReader struct {
	inner   Reader
	storage Storage
}

func NewStorageReader(inner Reader, storage Storage) *StorageReader {
	return &StorageReader{inner: inner, storage: storage}
}

func (this *StorageReader) Read() (interface{}, error) {
	if config, err := this.inner.Read(); err != nil {
		return nil, err
	} else {
		this.storage.Store(config)
		return config, nil
	}
}
