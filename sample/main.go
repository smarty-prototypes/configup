package main

import (
	"log"
	"time"

	"github.com/smartystreets/confighup"
)

func main() {
	jsonReader := NewJSONReader("/Users/jonathan/Code/src/github.com/smartystreets/confighup/sample/config.json")

	watcher := confighup.New(jsonReader)
	if err := watcher.Initialize(); err != nil {
		log.Fatalln("[ERROR] Unable to read configuration:", err)
	}

	go watcher.Listen()

	for {
		log.Println(watcher.Load())
		time.Sleep(time.Second * 1)
	}
}
