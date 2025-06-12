package switchboard

import (
    "fx-switchboard/step"
    "go.uber.org/fx"
)

func Module(cfg Config) fx.Option {
    return fx.Provide(
        func() Config { return cfg },
        func(lc fx.Lifecycle, inj fx.Injector, cfg Config) ([]step.Step, error) {
            var results []step.Step
            for _, item := range cfg.Steps {
                factory := step.LookupFactory(item.Type)
                if factory == nil {
                    return nil, errUnknownType(item.Type)
                }
                instance, err := factory.Build(item.Config, lc, inj)
                if err != nil {
                    return nil, err
                }
                results = append(results, instance)
            }
            return results, nil
        },
    )
}

func errUnknownType(t string) error {
    return fmt.Errorf("unknown step type: %s", t)
}
