package modules

type Module interface {
	Name() string
	Execute(input map[string]interface{}) (map[string]interface{}, error)
	Validate(config map[string]interface{}) error
}

type Manager struct {
	modules map[string]Module
}

func NewManager() *Manager {
	return &Manager{
		modules: make(map[string]Module),
	}
}

func (m *Manager) Register(name string, module Module) {
	m.modules[name] = module
}

func (m *Manager) Get(name string) (Module, bool) {
	module, ok := m.modules[name]
	return module, ok
}

func (m *Manager) List() []string {
	names := make([]string, 0, len(m.modules))
	for name := range m.modules {
		names = append(names, name)
	}
	return names
}
