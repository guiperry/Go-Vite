package modules

type ExampleModule struct{}

func (m *ExampleModule) Name() string {
	return "example"
}

func (m *ExampleModule) Execute(input map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{
		"status": "success",
		"data":   input,
	}, nil
}

func (m *ExampleModule) Validate(config map[string]interface{}) error {
	return nil
}

func LoadBuiltinModules(manager *Manager) {
	manager.Register("example", &ExampleModule{})
}
