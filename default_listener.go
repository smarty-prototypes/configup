package configup

import (
	"sync"

	"github.com/smartystreets/logging"
)

type DefaultListener struct {
	signaler    Signaler
	reader      Reader
	storage     Storage
	subscribers []chan<- interface{}
	lock        *sync.RWMutex
	logger      *logging.Logger
}

func NewListener(signaler Signaler, reader Reader, storage Storage) *DefaultListener {
	return &DefaultListener{signaler: signaler, reader: reader, storage: storage, lock: &sync.RWMutex{}}
}

func (this *DefaultListener) Initialize() error {
	_, err := this.reader.Read()
	return err
}

func (this *DefaultListener) Listen() {
	for notification := range this.signaler.Channel() {
		this.logger.Printf("[INFO] Received [%s] signal, reloading configuration...\n", notification)
		this.reload()
	}
}
func (this *DefaultListener) reload() {
	if updated, err := this.reader.Read(); err != nil {
		this.logger.Printf("[ERROR] Unable to reload configuration: [%s]\n", err)
	} else {
		this.logger.Println("[INFO] Configuration reloaded successfully.")
		this.notify(updated)
	}
}
func (this *DefaultListener) notify(updated interface{}) {
	this.lock.RLock()
	defer this.lock.RUnlock()

	for _, subscriber := range this.subscribers {
		subscriber <- updated
	}
}

func (this *DefaultListener) Load() interface{} {
	return this.storage.Load()
}

func (this *DefaultListener) Subscribe(subscribers ...chan<- interface{}) {
	this.lock.Lock()
	this.subscribers = append(this.subscribers, subscribers...)
	this.lock.Unlock()
}
func (this *DefaultListener) Unsubscribe(subscriber chan<- interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()

	for i := range this.subscribers {
		if subscriber != this.subscribers[i] {
			continue
		}

		this.subscribers = append(this.subscribers[:i], this.subscribers[i+1:]...)
		break
	}
}
func (this *DefaultListener) Close() error {
	this.lock.Lock()
	defer this.lock.Unlock()

	this.signaler.Close()

	for _, subscriber := range this.subscribers {
		close(subscriber)
	}

	this.subscribers = this.subscribers[0:0]
	return nil
}
