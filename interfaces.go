package configup

import "os"

type Reader interface {
	Read() (interface{}, error)
}

type Storage interface {
	Store(interface{})
	StorageReader
}

type StorageReader interface {
	Load() interface{}
}

type Signaler interface {
	Open(...os.Signal)
	Close()
	Channel() <-chan os.Signal
}
