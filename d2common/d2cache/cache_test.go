package d2cache

import (
	"testing"
)

func TestCacheInsert(t *testing.T) {
	cache := CreateCache(1)
	insertError := cache.Insert("A", "", 1)

	if insertError != nil {
		t.Fatalf("Cache insert resulted in unexpected error: %s", insertError)
	}
}

func TestCacheInsertWithinBudget(t *testing.T) {
	cache := CreateCache(1)
	insertError := cache.Insert("A", "", 2)

	if insertError != nil {
		t.Fatalf("Cache insert resulted in unexpected error: %s", insertError)
	}
}

func TestCacheInsertUpdatesWeight(t *testing.T) {
	cache := CreateCache(2)
	_ = cache.Insert("A", "", 1)
	_ = cache.Insert("B", "", 1)
	_ = cache.Insert("budget_exceeded", "", 1)

	if cache.GetWeight() != 2 {
		t.Fatal("Cache with budget 2 did not correctly set weight after evicting one of three nodes")
	}
}

func TestCacheInsertDuplicateRejected(t *testing.T) {
	cache := CreateCache(2)
	_ = cache.Insert("dupe", "", 1)
	dupeError := cache.Insert("dupe", "", 1)

	if dupeError == nil {
		t.Fatal("Cache insert of duplicate key did not result in any err")
	}
}

func TestCacheInsertEvictsLeastRecentlyUsed(t *testing.T) {
	cache := CreateCache(2)
	// with a budget of 2, inserting 3 keys should evict the last
	_ = cache.Insert("evicted", "", 1)
	_ = cache.Insert("A", "", 1)
	_ = cache.Insert("B", "", 1)

	_, foundEvicted := cache.Retrieve("evicted")
	if foundEvicted {
		t.Fatal("Cache insert did not trigger eviction after weight exceedance")
	}

	// double check that only 1 one was evicted and not any extra
	_, foundA := cache.Retrieve("A")
	_, foundB := cache.Retrieve("B")

	if !foundA || !foundB {
		t.Fatal("Cache insert evicted more than necessary")
	}
}

func TestCacheInsertEvictsLeastRecentlyRetrieved(t *testing.T) {
	cache := CreateCache(2)
	_ = cache.Insert("A", "", 1)
	_ = cache.Insert("evicted", "", 1)

	// retrieve the oldest node, promoting it head so it is not evicted
	cache.Retrieve("A")

	// insert once more, exceeding weight capacity
	_ = cache.Insert("B", "", 1)
	// now the least recently used key should be evicted
	_, foundEvicted := cache.Retrieve("evicted")
	if foundEvicted {
		t.Fatal("Cache insert did not evict least recently used after weight exceedance")
	}
}

func TestClear(t *testing.T) {
	cache := CreateCache(1)
	_ = cache.Insert("cleared", "", 1)
	cache.Clear()
	_, found := cache.Retrieve("cleared")

	if found {
		t.Fatal("Still able to retrieve nodes after cache was cleared")
	}
}
