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

func New(reader Reader) *Wireup {
	return &Wireup{
		reader:   reader,
		storage:  &atomic.Value{},
		signaler: NewSignaler(),
	}
}

func (this *Wireup) WithStorage(storage Storage) *Wireup {
	this.storage = storage
	return this
}

func (this *Wireup) WithSignaler(signaler Signaler) *Wireup {
	this.signaler = signaler
	return this
}

func (this *Wireup) WithSignal(signals ...os.Signal) *Wireup {
	this.signals = append(this.signals, signals...)
	return this
}

func (this *Wireup) WithNotify(subscribers ...chan<- interface{}) *Wireup {
	this.subscribers = append(this.subscribers, subscribers...)
	return this
}

func (this *Wireup) Initialize() (*DefaultListener, error) {
	listener := NewListener(this.signaler, this.reader, this.storage)
	if err := listener.Initialize(); err != nil {
		return nil, err
	}

	listener.Subscribe(this.subscribers...)
	this.signaler.Open(this.signals...)
	return listener, nil
}
