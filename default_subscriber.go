package configup

import (
	"os"
	"os/signal"
	"syscall"
)

type DefaultSubscriber struct {
	subscription chan os.Signal
}

func NewSubscriber() *DefaultSubscriber {
	return &DefaultSubscriber{subscription: make(chan os.Signal, 16)}
}

func (this *DefaultSubscriber) Subscribe(signals ...os.Signal) {
	if len(signals) == 0 {
		signals = append(signals, syscall.SIGHUP)
	}

	signal.Notify(this.subscription, signals...)
}
func (this *DefaultSubscriber) Unsubscribe() {
	signal.Stop(this.subscription)
	close(this.subscription)
}
func (this *DefaultSubscriber) Subscription() chan os.Signal {
	return this.subscription
}
