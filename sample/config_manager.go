package main

import (
	"log"

	"github.com/smartystreets/confighup"
)

type ConfigManager struct {
	storage confighup.Storage
}

func NewConfigManager(path string) *ConfigManager {
	jsonReader := NewJSONReader(path)

	storage, err := confighup.New(jsonReader).Initialize()
	if err != nil {
		log.Fatalln("[ERROR] Unable to read configuration:", err)
	}

	return &ConfigManager{storage: storage}
}

func (this *ConfigManager) Config() ConfigFile {
	return this.storage.Load().(ConfigFile)
}
