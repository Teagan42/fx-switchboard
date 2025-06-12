package switchboard

type StepConfig struct {
    Type   string
    Config map[string]any
}

type Config struct {
    Steps []StepConfig
}
