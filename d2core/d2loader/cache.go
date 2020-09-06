package d2loader

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2loader/asset"
)

const (
	errFmtNoEntry = "no cache entry for key `%s`"
)

// Entry represents an entry in the cache, which is just an asset
type Entry interface {
	asset.Asset
}

func newCache() *Cache {
	return &Cache{make(map[string]Entry)}
}

// Cache represents the files which have already been loaded, which can be quickly retrieved
// if/when the loader tries to load them again
type Cache struct {
	entries map[string]Entry
}

// Clear will remove all cache entries
func (c *Cache) Clear() {
	c.entries = make(map[string]Entry)
}

// Add will add the given cache entry with the given key
func (c *Cache) Add(key string, entry Entry) {
	c.entries[key] = entry
}

// Get will return a cache entry for the given key
func (c *Cache) Get(key string) (Entry, error) {
	data, found := c.entries[key]

	if !found {
		return nil, fmt.Errorf(errFmtNoEntry, key)
	}

	return data, nil
}
