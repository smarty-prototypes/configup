package confighup

import (
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
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

func (this *Wireup) WatchSignals(signals ...os.Signal) *Wireup {
	for _, item := range signals {
		this.signals = append(this.signals, item)
	}

	return this
}

func (this *Wireup) Initialize() (*Watcher, error) {
	reader := NewReader(this.reader, this.storage)
	if _, err := reader.Read(); err != nil {
		return nil, err
	}

	listener := NewListener(this.signalChannel(), reader)
	return &Watcher{listener: listener, storage: this.storage}, nil
}

func (this *Wireup) signalChannel() chan os.Signal {
	channel := make(chan os.Signal, 16)

	if len(this.signals) == 0 {
		signal.Notify(channel, syscall.SIGHUP)
	} else {
		signal.Notify(channel, this.signals...)
	}

	return channel
}
