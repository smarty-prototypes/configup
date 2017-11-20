package main

import "github.com/smartystreets/configup"

type Configuration struct {
	Username string
	Password string
}

type ConfigManager struct {
	storage configup.StorageReader
}

func NewConfigManager(path string) *ConfigManager {
	storage := configup.FromJSONFile(path, &Configuration{})
	return &ConfigManager{storage: storage}
}

func (this *ConfigManager) Config() *Configuration {
	// safely returns the latest instance of Configuration for use
	// by multiple goroutines.
	return this.storage.Load().(*Configuration)
}
