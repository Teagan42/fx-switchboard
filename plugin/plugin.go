package plugin

import "go.uber.org/fx"

type Plugin struct {
    Metadata
}

type plugin interface {
    Execute() error
}

type PluginFactory interface {
    Name() string
    Metadata() string
    Description() *string
    Build(cfg map[string]any, lc fx.Lifecycle, inj fx.Injector) (*Plugin, error)
}

func (f *PluginFactory) Name() string {
    return f.Metadata().Name
}

func (f *PluginFactory) Description() *string {
    return f.Metadata().Description
}