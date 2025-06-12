package step

var registry = make(map[string]Factory)

func Register(name string, factory Factory) {
    registry[name] = factory
}

func LookupFactory(name string) Factory {
    return registry[name]
}
