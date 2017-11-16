package main

import "github.com/smartystreets/configup"

type ConfigManager struct {
	storage configup.Storage
}

func NewConfigManager(path string) *ConfigManager {
	storage := configup.FromJSONFile(path, &ConfigFile{})
	return &ConfigManager{storage: storage}
}

func (this *ConfigManager) Config() *ConfigFile {
	return this.storage.Load().(*ConfigFile)
}
