package d2interface

// Cache stores arbitrary data for fast retrieval
type Cache interface {
	SetVerbose(verbose bool)
	GetWeight() int
	GetBudget() int
	Insert(key string, value interface{}, weight int) error
	Retrieve(key string) (interface{}, bool)
	Clear()
}

// Cacher is something that has a cache
type Cacher interface {
	ClearCache()
	GetCache() Cache
}
