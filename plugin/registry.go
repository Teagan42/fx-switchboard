package plugin

import (
    "go.uber.org/fx"
    "sync"
)

var (
    pluginFactories = map[string]PluginFactory{}
    mutex   sync.Mutex
)

type pluginRegistry interface {
    Register(factories []PluginFactory...) *pluginRegistry
    GetFactories() []PluginFactory
} 

func (pr *pluginRegistry) Register(factories []PluginFactory...) *pluginRegistry {
    mutex.Lock()
    defer mutex.Unlock()
    for _, factory := factories {
        pluginFactories[
    }
    pluginFactories = append(pluginFactories, factories...)
    return pr
}

func (pr *pluginRegistry) GetFactories() []PluginFactory {
    mutex.Lock()
    defer mutex.Unlock()
    return append([]PluginFactory(nil), pluginFactories...)
}

var PluginRegistry = pluginRegistry{}