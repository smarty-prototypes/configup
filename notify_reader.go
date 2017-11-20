package configup

import "sync"

type NotifyReader struct {
	inner       Reader
	subscribers []chan<- interface{}
	lock        *sync.RWMutex
}

func NewNotifyReader(reader Reader) *NotifyReader {
	return &NotifyReader{inner: reader, lock: &sync.RWMutex{}}
}

func (this *NotifyReader) Read() (interface{}, error) {
	if value, err := this.inner.Read(); err != nil {
		return nil, err
	} else {
		this.notify(value)
		return value, nil
	}
}
func (this *NotifyReader) notify(value interface{}) {
	this.lock.RLock()
	defer this.lock.RUnlock()

	for _, subscriber := range this.subscribers {
		subscriber <- value
	}

}

func (this *NotifyReader) Subscribe(subscribers ...chan<- interface{}) {
	this.lock.Lock()
	this.subscribers = append(this.subscribers, subscribers...)
	this.lock.Unlock()
}
func (this *NotifyReader) Unsubscribe(subscriber chan<- interface{}) {
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
func (this *NotifyReader) Close() error {
	this.lock.Lock()
	defer this.lock.Unlock()

	for _, subscriber := range this.subscribers {
		close(subscriber)
	}

	this.subscribers = this.subscribers[0:0]
	return nil
}
