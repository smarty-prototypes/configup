package configup

type DefaultReader struct {
	inner   Reader
	storage Storage
}

func NewReader(inner Reader, storage Storage) *DefaultReader {
	return &DefaultReader{inner: inner, storage: storage}
}

func (this *DefaultReader) Read() (interface{}, error) {
	if config, err := this.inner.Read(); err != nil {
		return nil, err
	} else {
		this.storage.Store(config)
		return config, nil
	}
}
