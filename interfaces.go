package configup

import (
	"io"
	"os"
)

type Reader interface {
	Read() (interface{}, error)
}

type Storage interface {
	Store(interface{})
	Load() interface{}
}

type Listener interface {
	Listen()
	io.Closer
}

type Subscriber interface {
	Subscribe(...os.Signal)
	Unsubscribe()
	Subscription() chan os.Signal
}
