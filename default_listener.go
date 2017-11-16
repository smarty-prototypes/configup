package confighup

import "github.com/smartystreets/logging"

type DefaultListener struct {
	subscriber Subscriber
	reader     Reader
	logger     *logging.Logger
}

func NewListener(subscriber Subscriber, reader Reader) *DefaultListener {
	return &DefaultListener{subscriber: subscriber, reader: reader}
}

func (this *DefaultListener) Listen() {
	for notification := range this.subscriber.Subscription() {
		this.logger.Printf("[INFO] Received [%s] signal, reloading configuration...\n", notification)
		this.reload()
	}
}

func (this *DefaultListener) reload() {
	if _, err := this.reader.Read(); err != nil {
		this.logger.Printf("[ERROR] Unable to reload configuration: [%s]\n", err)
	} else {
		this.logger.Println("[INFO] Configuration reloaded successfully.")
	}
}

func (this *DefaultListener) Close() error {
	this.subscriber.Unsubscribe()
	return nil
}
