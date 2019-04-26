package module

type Module struct {
	name string
	version string
}

func (m Module) Name() string {
	return m.name
}

func (m Module) Version() string {
	return m.version
}