package confighup

import (
	"os"

	"github.com/smartystreets/logging"
)

type Listener struct {
	signals chan os.Signal
	reader  Reader
	logger  *logging.Logger
}

func NewListener(signals chan os.Signal, reader Reader) *Listener {
	return &Listener{signals: signals, reader: reader}
}

func (this *Listener) Listen() {
	for notification := range this.signals {
		this.logger.Printf("[INFO] Received [%s] signal, reloading configuration...\n", notification)
		this.reload()
	}
}

func (this *Listener) reload() {
	if _, err := this.reader.Read(); err != nil {
		this.logger.Printf("[ERROR] Unable to reload configuration: [%s]\n", err)
	} else {
		this.logger.Println("[INFO] Configuration reloaded successfully.")
	}
}
