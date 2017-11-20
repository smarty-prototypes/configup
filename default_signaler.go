package configup

import (
	"os"
	"os/signal"
	"syscall"
)

type DefaultSignaler struct {
	channel chan os.Signal
}

func NewSignaler() *DefaultSignaler {
	return &DefaultSignaler{channel: make(chan os.Signal, 16)}
}

func (this *DefaultSignaler) Open(signals ...os.Signal) {
	if len(signals) == 0 {
		signals = append(signals, syscall.SIGHUP)
	}

	signal.Notify(this.channel, signals...)
}

func (this *DefaultSignaler) Close() {
	signal.Stop(this.channel)
	close(this.channel)
}

func (this *DefaultSignaler) Channel() <-chan os.Signal {
	return this.channel
}
