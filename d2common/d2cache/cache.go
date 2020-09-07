package d2cache

import (
	"errors"
	"log"
	"sync"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

var _ d2interface.Cache = &Cache{} // Static check to confirm struct conforms to interface

type cacheNode struct {
	next   *cacheNode
	prev   *cacheNode
	key    string
	value  interface{}
	weight int
}

// Cache stores arbitrary data for fast retrieval
type Cache struct {
	head    *cacheNode
	tail    *cacheNode
	lookup  map[string]*cacheNode
	weight  int
	budget  int
	verbose bool
	mutex   sync.Mutex
}

// CreateCache creates an  instance of a Cache
func CreateCache(budget int) d2interface.Cache {
	return &Cache{lookup: make(map[string]*cacheNode), budget: budget}
}

// SetVerbose turns on verbose printing (warnings and stuff)
func (c *Cache) SetVerbose(verbose bool) {
	c.verbose = verbose
}

// GetWeight gets the "weight" of a cache
func (c *Cache) GetWeight() int {
	return c.weight
}

// GetBudget gets the memory budget of a cache
func (c *Cache) GetBudget() int {
	return c.budget
}

// Insert inserts an object into the cache
func (c *Cache) Insert(key string, value interface{}, weight int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, found := c.lookup[key]; found {
		return errors.New("key already exists in Cache")
	}

	node := &cacheNode{
		key:    key,
		value:  value,
		weight: weight,
		next:   c.head,
	}

	if c.head != nil {
		c.head.prev = node
	}

	c.head = node
	if c.tail == nil {
		c.tail = node
	}

	c.lookup[key] = node
	c.weight += node.weight

	for ; c.tail != nil && c.tail != c.head && c.weight > c.budget; c.tail = c.tail.prev {
		c.weight -= c.tail.weight
		c.tail.prev.next = nil

		if c.verbose {
			log.Printf(
				"warning -- Cache is evicting %s (%d) for %s (%d); spare weight is now %d",
				c.tail.key,
				c.tail.weight,
				key,
				weight,
				c.budget-c.weight,
			)
		}

		delete(c.lookup, c.tail.key)
	}

	return nil
}

// Retrieve gets an object out of the cache
func (c *Cache) Retrieve(key string) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	node, found := c.lookup[key]
	if !found {
		return nil, false
	}

	if node != c.head {
		if node.next != nil {
			node.next.prev = node.prev
		}

		if node.prev != nil {
			node.prev.next = node.next
		}

		if node == c.tail {
			c.tail = c.tail.prev
		}

		node.next = c.head
		node.prev = nil

		if c.head != nil {
			c.head.prev = node
		}

		c.head = node
	}

	return node.value, true
}

// Clear removes all cache entries
func (c *Cache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.head = nil
	c.tail = nil
	c.lookup = make(map[string]*cacheNode)
	c.weight = 0
}
