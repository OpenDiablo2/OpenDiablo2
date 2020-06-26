package d2mapengine

type MapGeneratorWilderness struct {
	seed   int64
	level  *MapLevel
	engine *MapEngine
}

func (m *MapGeneratorWilderness) init(s int64, l *MapLevel, e *MapEngine) {
	m.seed = s
	m.level = l
	m.engine = e
}

func (m *MapGeneratorWilderness) generate() {}
