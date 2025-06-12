package step

import "go.uber.org/fx"

type Step interface {
    Execute() error
}

type Factory interface {
    Build(cfg map[string]any, lc fx.Lifecycle, inj fx.Injector) (Step, error)
}
