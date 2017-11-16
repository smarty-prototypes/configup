package main

import (
	"log"

	"github.com/smartystreets/confighup"
)

type Configuration struct {
	storage confighup.Storage
}

func NewConfiguration(configFile string) *Configuration {
	jsonReader := NewJSONReader(configFile)

	storage, err := confighup.New(jsonReader).Initialize()
	if err != nil {
		log.Fatalln("[ERROR] Unable to read configuration:", err)
	}

	return &Configuration{storage: storage}
}

func (this *Configuration) Values() ConfigValues {
	return this.storage.Load().(ConfigValues)
}
