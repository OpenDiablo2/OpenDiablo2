package d2mapengine

type MapGenerator interface {
	init(seed int64, level *MapLevel, engine *MapEngine)
	generate()
}
