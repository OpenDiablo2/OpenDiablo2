package d2loader

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2cache"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/asset/types"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/filesystem"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/mpq"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
)

const (
	defaultCacheBudget      = 1024 * 1024 * 512
	defaultCacheEntryWeight = 1
	errFmtFileNotFound      = "file not found: %s"
)

const (
	logPrefix = "File Loader"
)

const (
	fontToken  = d2resource.LanguageFontToken
	tableToken = d2resource.LanguageTableToken
)

// NewLoader creates a new loader
func NewLoader(l d2util.LogLevel) (*Loader, error) {
	loader := &Loader{}

	loader.Cache = d2cache.CreateCache(defaultCacheBudget)
	loader.Logger = d2util.NewLogger()

	loader.Logger.SetPrefix(logPrefix)
	loader.Logger.SetLevel(l)

	return loader, nil
}

// Loader represents the manager that handles loading and caching assets with the asset Sources
// that have been added
type Loader struct {
	language *string
	charset  *string
	d2interface.Cache
	*d2util.Logger
	Sources []asset.Source
}

// SetLanguage sets the language for loader
func (l *Loader) SetLanguage(language *string) {
	l.language = language
}

// SetCharset sets the charset for loader
func (l *Loader) SetCharset(charset *string) {
	l.charset = charset
}

// Load attempts to load an asset with the given sub-path. The sub-path is relative to the root
// of each asset source root (regardless of the type of asset source)
func (l *Loader) Load(subPath string) (asset.Asset, error) {
	subPath = filepath.Clean(subPath)

	if l.language != nil {
		charset := l.charset
		language := l.language

		subPath = strings.ReplaceAll(subPath, fontToken, *charset)
		subPath = strings.ReplaceAll(subPath, tableToken, *language)
	}

	// first, we check the cache for an existing entry
	if cached, found := l.Retrieve(subPath); found {
		l.Debug(fmt.Sprintf("Retrieved `%s` from cache", subPath))

		a := cached.(asset.Asset)
		_, err := a.Seek(0, 0)

		return a, err
	}

	// if it isn't in the cache, we check if each source can open the file
	for idx := range l.Sources {
		source := l.Sources[idx]

		// if the source can open the file, then we cache it and return it
		loadedAsset, err := source.Open(subPath)
		if err != nil {
			l.Debug(fmt.Sprintf("Checked `%s`, file not found", source.Path()))
			continue
		}

		srcBase, _ := filepath.Abs(source.Path())
		l.Info(fmt.Sprintf("from %s, loading %s", srcBase, subPath))

		return loadedAsset, l.Insert(subPath, loadedAsset, defaultCacheEntryWeight)
	}

	return nil, fmt.Errorf(errFmtFileNotFound, subPath)
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
		l.Error(err.Error())
		return nil, err
	}

	mode := info.Mode()

	sourceType := types.AssetSourceUnknown

	if mode.IsDir() {
		sourceType = types.AssetSourceFileSystem
	}

	if mode.IsRegular() {
		sourceType = types.CheckSourceType(cleanPath)
	}

	switch sourceType {
	case types.AssetSourceMPQ:
		source, err := mpq.NewSource(cleanPath)
		if err == nil {
			l.Info(fmt.Sprintf("adding MPQ source `%s`", cleanPath))
			l.Sources = append(l.Sources, source)

			return source, nil
		}
	case types.AssetSourceFileSystem:
		source := &filesystem.Source{
			Root: cleanPath,
		}

		l.Info(fmt.Sprintf("adding filesystem source `%s`", cleanPath))
		l.Sources = append(l.Sources, source)

		return source, nil
	case types.AssetSourceUnknown:
		l.Warning(fmt.Sprintf("unknown asset source `%s`", cleanPath))
	}

	return nil, fmt.Errorf("unknown asset source `%s`", cleanPath)
}
