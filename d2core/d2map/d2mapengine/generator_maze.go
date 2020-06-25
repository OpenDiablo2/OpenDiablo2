package d2mapengine

type MapGeneratorMaze struct {
	seed   int64
	level  *MapLevel
	engine *MapEngine
}

func (m *MapGeneratorMaze) init(s int64, l *MapLevel, e *MapEngine) {
	m.seed = s
	m.level = l
	m.engine = e
}

func (m *MapGeneratorMaze) generate() {}
