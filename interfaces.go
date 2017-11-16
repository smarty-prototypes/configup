package confighup

type Reader interface {
	Read() (interface{}, error)
}

type Storage interface {
	Store(interface{})
	Load() interface{}
}
