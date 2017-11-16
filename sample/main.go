package main

import (
	"log"
	"os"
	"path"
	"time"
)

func main() {
	configFile, _ := os.Getwd()
	configFile = path.Join(configFile, "config.json")
	manager := NewConfigManager(configFile)

	for {
		log.Println(manager.Config())
		time.Sleep(time.Second * 1)
	}
}
