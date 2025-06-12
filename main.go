package main

import (
    "fx-switchboard/plugin"
    "fx-switchboard/plugins/shell"
    "fx-switchboard/switchboard"
    "fx-switchboard/step"
    "go.uber.org/fx"
    "go.uber.org/zap"
)

func run(steps []step.Step) {
    for _, s := range steps {
        s.Execute()
    }
}

func main() {
    cfg := switchboard.Config{
        Steps: []switchboard.StepConfig{
            {Type: "shell", Config: map[string]any{"script": "echo Hello Plugins"}},
        },
    }

    app := fx.New(
        fx.Provide(func() *zap.Logger {
            return zap.NewExample()
        }),
        fx.Options(plugin.Modules()...),
        switchboard.Module(cfg),
        shell.Module(),
        fx.Invoke(run),
    )

    app.Run()
}
