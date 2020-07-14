package d2asset

import (
	"errors"
	"path"
	"sync"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2mpq"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
)

// Static checks to confirm struct conforms to interface
var _ d2interface.ArchiveManager = &archiveManager{}
var _ d2interface.Cacher = &archiveManager{}

type archiveManager struct {
	cache    d2interface.Cache
	config   *d2config.Configuration
	archives []d2interface.Archive
	mutex    sync.Mutex
}

const (
	archiveBudget = 1024 * 1024 * 512
)

func createArchiveManager(config *d2config.Configuration) d2interface.ArchiveManager {
	return &archiveManager{cache: d2common.CreateCache(archiveBudget), config: config}
}

// LoadArchiveForFile loads the archive for the given (in-archive) file path
func (am *archiveManager) LoadArchiveForFile(filePath string) (d2interface.Archive, error) {
	am.mutex.Lock()
	defer am.mutex.Unlock()

	if err := am.CacheArchiveEntries(); err != nil {
		return nil, err
	}

	for _, archive := range am.archives {
		if archive.Contains(filePath) {
			result, ok := am.LoadArchive(archive.Path())
			if ok == nil {
				return result, nil
			}
		}
	}

	return nil, errors.New("file not found")
}

// FileExistsInArchive checks if a file exists in an archive
func (am *archiveManager) FileExistsInArchive(filePath string) (bool, error) {
	am.mutex.Lock()
	defer am.mutex.Unlock()

	if err := am.CacheArchiveEntries(); err != nil {
		return false, err
	}

	for _, archiveEntry := range am.archives {
		if archiveEntry.Contains(filePath) {
			return true, nil
		}
	}

	return false, nil
}

// LoadArchive loads and caches an archive
func (am *archiveManager) LoadArchive(archivePath string) (d2interface.Archive, error) {
	if archive, found := am.cache.Retrieve(archivePath); found {
		return archive.(d2interface.Archive), nil
	}

	archive, err := d2mpq.Load(archivePath)
	if err != nil {
		return nil, err
	}

	if err := am.cache.Insert(archivePath, archive, int(archive.Size())); err != nil {
		return nil, err
	}

	return archive, nil
}

// CacheArchiveEntries updates the archive entries
func (am *archiveManager) CacheArchiveEntries() error {
	if len(am.archives) == len(am.config.MpqLoadOrder) {
		return nil
	}

	am.archives = nil

	for _, archiveName := range am.config.MpqLoadOrder {
		archivePath := path.Join(am.config.MpqPath, archiveName)

		archive, err := am.LoadArchive(archivePath)
		if err != nil {
			return err
		}

		am.archives = append(
			am.archives,
			archive,
		)
	}

	return nil
}

// ClearCache clears the archive manager cache
func (am *archiveManager) ClearCache() {
	am.cache.Clear()
}

// GetCache returns the archive manager cache
func (am *archiveManager) GetCache() d2interface.Cache {
	return am.cache
}
