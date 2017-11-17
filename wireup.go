package configup

import (
	"os"
	"sync/atomic"
)

type Wireup struct {
	reader      Reader
	storage     Storage
	signals     []os.Signal
	subscribers []chan interface{}
}

func New(reader Reader) *Wireup {
	return &Wireup{
		reader:  reader,
		storage: &atomic.Value{},
	}
}

func (this *Wireup) WithSignal(signals ...os.Signal) *Wireup {
	this.signals = append(this.signals, signals...)
	return this
}

func (this *Wireup) WithNotify(subscribers ...chan interface{}) *Wireup {
	this.subscribers = append(this.subscribers, subscribers...)
	return this
}

func (this *Wireup) Initialize() (Storage, error) {
	var reader Reader

	reader = NewReader(this.reader, this.storage)
	if _, err := reader.Read(); err != nil {
		return nil, err
	}

	if len(this.subscribers) > 0 {
		reader = NewNotifyReader(reader).Subscribe(this.subscribers...)
	}

	subscriber := NewSubscriber()
	subscriber.Subscribe(this.signals...)

	listener := NewListener(subscriber, reader)
	go listener.Listen()

	return this.storage, nil
}
