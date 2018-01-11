package configup

import (
	"os"
	"sync/atomic"
)

type Wireup struct {
	reader      Reader
	storage     Storage
	signaler    Signaler
	signals     []os.Signal
	subscribers []chan<- interface{}
}

func New(reader Reader, options ...Option) (*DefaultListener, error) {
	wireup := &Wireup{
		reader:   reader,
		storage:  &atomic.Value{},
		signaler: NewSignaler(),
	}
	for _, option := range options {
		option(wireup)
	}
	return wireup.initialize()
}

func (this *Wireup) initialize() (*DefaultListener, error) {
	listener := NewListener(this.signaler, this.reader, this.storage)
	if err := listener.Initialize(); err != nil {
		return nil, err
	}

	listener.Subscribe(this.subscribers...)
	this.signaler.Open(this.signals...)
	return listener, nil
}

type Option func(*Wireup)

func WithStorage(storage Storage) Option {
	return func(wireup *Wireup) { wireup.storage = storage }
}
func WithSignaler(signaler Signaler) Option {
	return func(wireup *Wireup) { wireup.signaler = signaler }
}
func WithSignal(signals ...os.Signal) Option {
	return func(wireup *Wireup) { wireup.signals = append(wireup.signals, signals...) }
}

func WithNotify(subscribers ...chan<- interface{}) Option {
	return func(wireup *Wireup) { wireup.subscribers = append(wireup.subscribers, subscribers...) }
}
