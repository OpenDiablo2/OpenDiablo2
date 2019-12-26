package d2asset

import (
	"errors"
	"log"
	"sync"
)

type cacheNode struct {
	next   *cacheNode
	prev   *cacheNode
	key    string
	value  interface{}
	weight int
}

type cache struct {
	head    *cacheNode
	tail    *cacheNode
	lookup  map[string]*cacheNode
	weight  int
	budget  int
	verbose bool
	mutex   sync.Mutex
}

func createCache(budget int) *cache {
	return &cache{lookup: make(map[string]*cacheNode), budget: budget}
}

func (c *cache) insert(key string, value interface{}, weight int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, found := c.lookup[key]; found {
		return errors.New("key already exists in cache")
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
			log.Printf("warning, cache evicting %s (%d) for %s (%d)", c.tail.key, c.tail.weight, key, weight)
		}

		delete(c.lookup, c.tail.key)
	}

	return nil
}

func (c *cache) retrieve(key string) (interface{}, bool) {
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

func (c *cache) clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.head = nil
	c.tail = nil
	c.lookup = make(map[string]*cacheNode)
	c.weight = 0
}
