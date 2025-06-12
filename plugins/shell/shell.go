package shell

import (
    "fmt"
    "fx-switchboard/plugin"
    "fx-switchboard/step"
    "go.uber.org/fx"
    "go.uber.org/zap"
)

type ShellPlugin struct {
    Script string
    Logger *zap.Logger
}

func (s *ShellPlugin) Execute() error {
    s.Logger.Info("ShellStep", zap.String("script", s.Script))
    fmt.Println(">>", s.Script)
    return nil
}

type Deps struct {
    fx.In
    Logger *zap.Logger
}

type ShellPluginFactory struct {
    Deps
}

func (f ShellPluginFactory) Build(cfg map[string]any, _ fx.Lifecycle, _ fx.Injector) (step.Step, error) {
    script, ok := cfg["script"].(string)
    if !ok {
        return nil, fmt.Errorf("missing or invalid 'script'")
    }
    return &ShellStep{Script: script, Logger: f.Logger}, nil
}

func Module() fx.Option {
    return fx.Provide(func(d Deps) step.Factory {
        f := ShellFactory{d}
        step.Register("shell", f)
        plugin.Register("shell", fx.Provide(func() step.Factory { return f }))
        return f
    })
}
