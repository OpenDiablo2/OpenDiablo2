package d2loader

import (
	"errors"
	"fmt"

	"os"
	"path/filepath"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2cache"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/asset/types"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/filesystem"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/mpq"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
)

const (
	defaultCacheBudget      = 1024 * 1024 * 512
	defaultCacheEntryWeight = 1
	errFileNotFound         = "file not found"
)

// NewLoader creates a new loader
func NewLoader() *Loader {
	loader := &Loader{}
	loader.Cache = d2cache.CreateCache(defaultCacheBudget)

	return loader
}

// Loader represents the manager that handles loading and caching assets with the asset Sources
// that have been added
type Loader struct {
	d2interface.Cache
	*d2util.Logger
	sources []asset.Source
}

// Load attempts to load an asset with the given sub-path. The sub-path is relative to the root
// of each asset source root (regardless of the type of asset source)
func (l *Loader) Load(subPath string) (asset.Asset, error) {
	subPath = filepath.Clean(subPath)

	// first, we check the cache for an existing entry
	if cached, found := l.Retrieve(subPath); found {
		l.Debug(fmt.Sprintf("file `%s` exists in loader cache", subPath))
		return cached.(asset.Asset), nil
	}

	// if it isn't in the cache, we check if each source can open the file
	for idx := range l.Sources {
		source := l.Sources[idx]

		// if the source can open the file, then we cache it and return it
		if loadedAsset, err := source.Open(subPath); err == nil {
			err := l.Insert(subPath, loadedAsset, defaultCacheEntryWeight)
			return loadedAsset, err
		}
	}

	return nil, errors.New(errFileNotFound)
}

// AddSource adds an asset source with the given path. The path will either resolve to a directory
// or a file on the host filesystem. In the case that it is a file, the file extension is used
// to determine the type of asset source. In the case that the path points to a directory, a
// FileSystemSource will be added.
func (l *Loader) AddSource(path string) (asset.Source, error) {
	if l.Sources == nil {
		l.Sources = make([]asset.Source, 0)
	}

	cleanPath := filepath.Clean(path)

	info, err := os.Lstat(cleanPath)
	if err != nil {
		l.Warning(err.Error())
		return nil, err
	}

	mode := info.Mode()

	sourceType := types.AssetSourceUnknown

	if mode.IsDir() {
		sourceType = types.AssetSourceFileSystem
	}

	if mode.IsRegular() {
		ext := filepath.Ext(cleanPath)
		sourceType = types.Ext2SourceType(ext)
	}

	switch sourceType {
	case types.AssetSourceMPQ:
		source, err := mpq.NewSource(cleanPath)
		if err == nil {
			l.Debug(fmt.Sprintf("adding MPQ source `%s`", cleanPath))
			l.Sources = append(l.Sources, source)

			return source, nil
		}
	case types.AssetSourceFileSystem:
		source := &filesystem.Source{
			Root: cleanPath,
		}

		l.Debug(fmt.Sprintf("adding filesystem source `%s`", cleanPath))
		l.Sources = append(l.Sources, source)

		return source, nil
	case types.AssetSourceUnknown:
		l.Warning(fmt.Sprintf("unknown asset source `%s`", cleanPath))
	}

	return nil, fmt.Errorf("unknown asset source `%s`", cleanPath)
}
