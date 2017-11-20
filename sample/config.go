package main

import "github.com/smartystreets/configup"

type Configuration struct {
	Username string
	Password string
}

type ConfigManager struct {
	listener      *configup.DefaultListener
	notifications chan interface{}
}

func NewConfigManager(path string) *ConfigManager {
	listener := configup.FromJSONFile(path, &Configuration{})
	notifications := make(chan interface{}, 16)
	listener.Subscribe(notifications)
	return &ConfigManager{listener: listener, notifications: notifications}
}

func (this *ConfigManager) Notifications() *Configuration {
	updated := <-this.notifications
	return updated.(*Configuration)
}

func (this *ConfigManager) Config() *Configuration {
	// safely returns the latest instance of Configuration for use
	// by multiple goroutines.
	return this.listener.Load().(*Configuration)
}
