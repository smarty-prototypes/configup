package main

import (
	"log"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"

	"github.com/smartystreets/confighup"
)

func main() {
	jsonReader := NewJSONReader("/Users/jonathan/Code/src/github.com/smartystreets/confighup/sample/config.json")

	storage := &atomic.Value{}
	storageReader := confighup.NewStorageReader(jsonReader, storage)

	subscription := make(chan interface{}, 16)
	notifyReader := confighup.NewNotifyReader(storageReader)
	notifyReader.Subscribe(subscription)

	go func() {
		log.Println("[INFO] Listening for notification")

		for change := range subscription {
			log.Printf("[INFO] Loaded config (1): %#v\n", change)
		}
	}()

	if _, err := notifyReader.Read(); err != nil {
		log.Println("[ERROR] Unable to read configuration:", err)
		return
	}

	signals := make(chan os.Signal, 16)
	signal.Notify(signals, syscall.SIGHUP)
	listener := confighup.NewListener(signals, notifyReader)

	listener.Listen()
}
