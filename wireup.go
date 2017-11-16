package configup

import (
	"log"
	"os"
	"sync/atomic"
)

type Wireup struct {
	reader  Reader
	storage Storage
	signals []os.Signal
}

// FromJSONFile receives a filename and an instance of the type of struct
// into which the contents of the specified file should be unmarshalled.
// The instance is only used to derive type information. Rather than fill
// the provided instance, the Storage will give back a fresh copy of the
// unmarshalled data.
func FromJSONFile(filename string, instance interface{}) Storage {
	reader := NewJSONReader(filename, instance)
	return FromReader(reader)
}

func FromReader(reader Reader) Storage {
	if storage, err := New(reader).Initialize(); err != nil {
		log.Fatalln("[ERROR] Unable to read configuration:", err)
		return nil
	} else {
		return storage
	}
}

func New(reader Reader) *Wireup {
	return &Wireup{
		reader:  reader,
		storage: &atomic.Value{},
	}
}

func (this *Wireup) WatchSignals(signals ...os.Signal) *Wireup {
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
