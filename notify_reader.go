package configup

import "sync"

type NotifyReader struct {
	inner       Reader
	subscribers []chan interface{}
	lock        *sync.RWMutex
}

func NewNotifyReader(inner Reader) *NotifyReader {
	return &NotifyReader{inner: inner, lock: &sync.RWMutex{}}
}

func (this *NotifyReader) Subscribe(subscribers ...chan interface{}) Reader {
	this.lock.Lock()
	defer this.lock.Unlock()

	for _, subscriber := range subscribers {
		if subscriber != nil {
			this.subscribers = append(this.subscribers, subscriber)
		}
	}

	return this
}

func (this *NotifyReader) Read() (interface{}, error) {
	if config, err := this.inner.Read(); err != nil {
		return nil, err
	} else {
		this.notify(config)
		return config, nil
	}
}

func (this *NotifyReader) notify(config interface{}) {
	this.lock.RLock()
	defer this.lock.RUnlock()

	for _, subscriber := range this.subscribers {
		subscriber <- config
	}
}

func (this *NotifyReader) Close() error {
	this.lock.Lock()
	defer this.lock.Unlock()

	for _, subscriber := range this.subscribers {
		close(subscriber)
	}

	this.subscribers = this.subscribers[0:0]
	return nil
}
