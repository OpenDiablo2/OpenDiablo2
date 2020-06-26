package d2mapengine

type MapGeneratorPreset struct {
	seed   int64
	level  *MapLevel
	engine *MapEngine
}

func (m *MapGeneratorPreset) init(s int64, l *MapLevel, e *MapEngine) {
	m.seed = s
	m.level = l
	m.engine = e
}

func (m *MapGeneratorPreset) generate() {}
