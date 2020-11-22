package d2systems

import (
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2cache"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

const (
	languageTokenFont        = "{LANG_FONT}" // nolint:gosec // no security issue here...
	languageTokenStringTable = "{LANG}"
)

const (
	fileHandleCacheBudget      = 1024
	fileHandleCacheEntryWeight = 1
)

const (
	logPrefixFileHandleResolver = "File Handle Resolver"
)

// NewFileHandleResolver creates a new file handle resolver system
func NewFileHandleResolver() *FileHandleResolutionSystem {
	// this filter is for entities that have a file path and file type but no file handle.
	filesToSource := akara.NewFilter().
		Require(d2components.FilePath).
		Require(d2components.FileType).
		Forbid(d2components.FileHandle).
		Forbid(d2components.FileSource).
		Build()

	sourcesToUse := akara.NewFilter().
		RequireOne(d2components.FileSource).
		Build()

	fhr := &FileHandleResolutionSystem{
		BaseSubscriberSystem: akara.NewBaseSubscriberSystem(filesToSource, sourcesToUse),
		cache:                d2cache.CreateCache(fileHandleCacheBudget).(*d2cache.Cache),
		Logger:               d2util.NewLogger(),
	}

	fhr.SetPrefix(logPrefixFileHandleResolver)

	return fhr
}

// FileHandleResolutionSystem is responsible for using file sources to resolve files.
// File sources are checked in the order that the sources were added.
//
// A file source can be something like an MPQ archive or a file system directory on the host machine.
//
// A file handle is a primitive representation of  a loaded file; something that has data
// in the form of a byte slice, but has not been parsed into a more meaningful struct, like a DC6 animation.
type FileHandleResolutionSystem struct {
	*akara.BaseSubscriberSystem
	*d2util.Logger
	cache        *d2cache.Cache
	filesToLoad  *akara.Subscription
	sourcesToUse *akara.Subscription
	filePaths    *d2components.FilePathMap
	fileTypes    *d2components.FileTypeMap
	fileSources  *d2components.FileSourceMap
	fileHandles  *d2components.FileHandleMap
}

// Init initializes the system with the given world
func (m *FileHandleResolutionSystem) Init(_ *akara.World) {
	m.Info("initializing ...")

	m.filesToLoad = m.Subscriptions[0]
	m.sourcesToUse = m.Subscriptions[1]

	// try to inject the components we require, then cast the returned
	// abstract ComponentMap back to the concrete implementation
	m.filePaths = m.InjectMap(d2components.FilePath).(*d2components.FilePathMap)
	m.fileTypes = m.InjectMap(d2components.FileType).(*d2components.FileTypeMap)
	m.fileHandles = m.InjectMap(d2components.FileHandle).(*d2components.FileHandleMap)
	m.fileSources = m.InjectMap(d2components.FileSource).(*d2components.FileSourceMap)
}

// Update iterates over entities which have not had a file handle resolved.
// For each source, it attempts to load this file with the given source.
// If the file can be opened by the source, we create the file handle using that source.
func (m *FileHandleResolutionSystem) Update() {
	filesToLoad := m.filesToLoad.GetEntities()
	sourcesToUse := m.sourcesToUse.GetEntities()

	for _, fileID := range filesToLoad {
		for _, sourceID := range sourcesToUse {
			if m.loadFileWithSource(fileID, sourceID) {
				break
			}
		}
	}
}

// try to load a file with a source, returns true if loaded
func (m *FileHandleResolutionSystem) loadFileWithSource(fileID, sourceID akara.EID) bool {
	fp, found := m.filePaths.GetFilePath(fileID)
	if !found {
		return false
	}

	ft, found := m.fileTypes.GetFileType(fileID)
	if !found {
		return false
	}

	source, found := m.fileSources.GetFileSource(sourceID)
	if !found {
		return false
	}

	sourceFp, found := m.filePaths.GetFilePath(sourceID)
	if !found {
		return false
	}

	// replace the locale tokens if present
	if strings.Contains(fp.Path, languageTokenFont) {
		fp.Path = strings.ReplaceAll(fp.Path, languageTokenFont, "latin")
	} else if strings.Contains(fp.Path, languageTokenStringTable) {
		fp.Path = strings.ReplaceAll(fp.Path, languageTokenStringTable, "ENG")
	}

	cacheKey := m.makeCacheKey(fp.Path, sourceFp.Path)
	if entry, found := m.cache.Retrieve(cacheKey); found {
		component := m.fileHandles.AddFileHandle(fileID)
		component.Data = entry.(d2interface.DataStream)

		return true
	}

	data, err := source.Open(fp)
	if err != nil {
		// HACK: sound environment stuff doesnt specify the path, just the filename
		// so we gotta check this edge case
		if ft.Type != d2enum.FileTypeWAV {
			return false
		}

		if !strings.Contains(fp.Path, "sfx") {
			return false
		}

		tryPath := strings.ReplaceAll(fp.Path, "sfx", "music")
		tmpComponent := &d2components.FilePathComponent{Path: tryPath}

		cacheKey = m.makeCacheKey(tryPath, sourceFp.Path)
		if entry, found := m.cache.Retrieve(cacheKey); found {
			component := m.fileHandles.AddFileHandle(fileID)
			component.Data = entry.(d2interface.DataStream)
			fp.Path = tryPath

			return true
		}

		data, err = source.Open(tmpComponent)
		if err != nil {
			return false
		}

		fp.Path = tryPath
	}

	m.Infof("resolved `%s` with source `%s`", fp.Path, sourceFp.Path)

	component := m.fileHandles.AddFileHandle(fileID)
	component.Data = data

	if err := m.cache.Insert(cacheKey, data, fileHandleCacheEntryWeight); err != nil {
		m.Error(err.Error())
	}

	return true
}

func (m *FileHandleResolutionSystem) makeCacheKey(path, source string) string {
	const sep = "->"
	return strings.Join([]string{source, path}, sep)
}
