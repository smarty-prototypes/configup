package confighup

import (
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
)

type Watcher struct {
	storage  Storage
	reader   Reader
	listener Listener
}

func New(inner Reader, signals ...os.Signal) *Watcher {
	storage := &atomic.Value{}
	reader := NewReader(inner, storage)

	if len(signals) == 0 {
		signals = append(signals, syscall.SIGHUP)
	}

	channel := make(chan os.Signal, 16)
	signal.Notify(channel, signals...)
	listener := NewListener(channel, reader)

	return &Watcher{
		storage:  storage,
		reader:   reader,
		listener: listener,
	}
}

func (this *Watcher) Initialize() error {
	if _, err := this.reader.Read(); err != nil {
		return err
	} else {
		return nil
	}
}

func (this *Watcher) Load() interface{} {
	return this.storage.Load()
}

func (this *Watcher) Listen() {
	this.listener.Listen()
}
