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
	config := NewConfiguration(configFile)

	for {
		log.Println(config.Values())
		time.Sleep(time.Second * 1)
	}
}
