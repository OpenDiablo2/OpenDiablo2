package d2loader

import (
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2loader/asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2loader/asset/types"
)

func TestCache_newCache(t *testing.T) {
	c := newCache()

	if c.entries == nil {
		t.Error("cache entry map should be initialized when newCache is called")
	}
}

func TestCache_Add(t *testing.T) {
	c := newCache()
	entry := &mockAsset{}

	c.Add("a", entry)
	c.Add("b", entry)
	c.Add("c", entry)

	if len(c.entries) != 3 {
		t.Error("unexpected cache entry count after adding a cache entry")
	}
}

func TestCache_Clear(t *testing.T) {
	c := newCache()
	entry := &mockAsset{}

	c.Add("a", entry)
	c.Add("b", entry)
	c.Add("c", entry)

	c.Clear()

	if len(c.entries) != 0 {
		t.Error("unexpected cache entry count after clearing the cache")
	}
}

func TestCache_Get(t *testing.T) {
	c := newCache()
	entryA, entryB := &mockAsset{}, &mockAsset{}
	keyA, keyB := "a", "b"

	c.Add(keyA, entryA)
	c.Add(keyB, entryB)

	retrieveA, errA := c.Get(keyA)
	if errA != nil {
		t.Error(errA)
	}

	retrieveA2, errA := c.Get(keyA)
	if errA != nil {
		t.Error(errA)
	}

	retrieveB, errB := c.Get(keyB)
	if errB != nil {
		t.Error(errB)
	}

	if len(c.entries) != 2 {
		t.Error("unexpected cache entry count after adding a cache entry")
	}

	if retrieveA == retrieveB {
		t.Error("unexpected entry retrieved from cache")
	}

	if retrieveA != retrieveA2 {
		t.Error("unexpected entry retrieved from cache")
	}
}

type mockAsset struct {
	path string // nolint:structcheck // causes equality check issues if removed
}

func (m *mockAsset) Type() types.AssetType {
	return types.AssetTypeUnknown
}

func (m *mockAsset) Source() asset.Source {
	return nil
}

func (m *mockAsset) Path() string {
	return m.path
}

func (m *mockAsset) Read(_ []byte) (n int, err error) {
	return 0, nil
}

func (m *mockAsset) Seek(_ int64, _ int) (n int64, err error) {
	return 0, nil
}
