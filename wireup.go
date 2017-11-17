package configup

import (
	"os"
	"sync/atomic"
)

type Wireup struct {
	reader  Reader
	storage Storage
	signals []os.Signal
}

func New(reader Reader) *Wireup {
	return &Wireup{
		reader:  reader,
		storage: &atomic.Value{},
	}
}

func (this *Wireup) WithSignal(signals ...os.Signal) *Wireup {
	for _, item := range signals {
		this.signals = append(this.signals, item)
	}

	return this
}

func (this *Wireup) Initialize() (Storage, error) {
	reader := NewReader(this.reader, this.storage)
	if _, err := reader.Read(); err != nil {
		return nil, err
	}

	subscriber := NewSubscriber()
	subscriber.Subscribe(this.signals...)

	listener := NewListener(subscriber, reader)
	go listener.Listen()

	return this.storage, nil
}
