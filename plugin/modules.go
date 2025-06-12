package plugin

import (
    "go.uber.org/fx"
    "sync"
)

var (
    modules []fx.Option
    mutex     sync.Mutex
)


type PluginModuleRegistry interface {
    GetModules()    [].fx.Option
    Register(name string, module fx.Option) *PluginModuleRegistry
}

func (pmr *PluginModuleRegistry) Register(name string, module fx.Option) *PluginModuleRegistry {
    mutex.Lock()
    defer mutex.Unlock()
    modules = append(modules, module)
    return pmr
}

func (pmr *PluginModuleRegistry) GetModules() []fx.Option {
    mutex.Lock()
    defer mutex.Unlock()
    return append([]fx.Option(nil), modules...)
}