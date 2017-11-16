package main

import "github.com/smartystreets/confighup"

type ConfigManager struct {
	storage confighup.Storage
}

func NewConfigManager(path string) *ConfigManager {
	reader := NewJSONReader(path)
	storage := confighup.FromReader(reader)
	return &ConfigManager{storage: storage}
}

func (this *ConfigManager) Config() ConfigFile {
	return this.storage.Load().(ConfigFile)
}
