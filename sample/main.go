package main

import (
	"log"
	"time"

	"github.com/smartystreets/confighup"
)

func main() {
	jsonReader := NewJSONReader(configFile)

	watcher, err := confighup.New(jsonReader).Initialize()
	if err != nil {
		log.Fatalln("[ERROR] Unable to read configuration:", err)
	}

	go watcher.Listen()

	for {
		log.Println(watcher.Load())
		time.Sleep(time.Second * 1)
	}
}

const configFile = "/Users/jonathan/Code/src/github.com/smartystreets/confighup/sample/config.json"
