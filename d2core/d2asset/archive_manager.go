package d2asset

import (
	"errors"
	"path"
	"sync"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2mpq"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
)

type archiveEntry struct {
	archivePath  string
	hashEntryMap d2mpq.HashEntryMap
}

type archiveManager struct {
	cache   *d2common.Cache
	config  *d2config.Configuration
	entries []archiveEntry
	mutex   sync.Mutex
}

const (
	archiveBudget = 1024 * 1024 * 512
)

func createArchiveManager(config *d2config.Configuration) *archiveManager {
	return &archiveManager{cache: d2common.CreateCache(archiveBudget), config: config}
}

func (am *archiveManager) loadArchiveForFile(filePath string) (*d2mpq.MPQ, error) {
	am.mutex.Lock()
	defer am.mutex.Unlock()

	if err := am.cacheArchiveEntries(); err != nil {
		return nil, err
	}

	for _, archiveEntry := range am.entries {
		if archiveEntry.hashEntryMap.Contains(filePath) {
			result, ok := am.loadArchive(archiveEntry.archivePath)
			if ok == nil {
				return result, nil
			}
		}
	}

	return nil, errors.New("file not found")
}

func (am *archiveManager) fileExistsInArchive(filePath string) (bool, error) {
	am.mutex.Lock()
	defer am.mutex.Unlock()

	if err := am.cacheArchiveEntries(); err != nil {
		return false, err
	}

	for _, archiveEntry := range am.entries {
		if archiveEntry.hashEntryMap.Contains(filePath) {
			return true, nil
		}
	}

	return false, nil
}

func (am *archiveManager) loadArchive(archivePath string) (*d2mpq.MPQ, error) {
	if archive, found := am.cache.Retrieve(archivePath); found {
		return archive.(*d2mpq.MPQ), nil
	}

	archive, err := d2mpq.Load(archivePath)
	if err != nil {
		return nil, err
	}

	if err := am.cache.Insert(archivePath, archive, int(archive.Data.ArchiveSize)); err != nil {
		return nil, err
	}

	return archive, nil
}

func (am *archiveManager) cacheArchiveEntries() error {
	if len(am.entries) == len(am.config.MpqLoadOrder) {
		return nil
	}

	am.entries = nil

	for _, archiveName := range am.config.MpqLoadOrder {
		archivePath := path.Join(am.config.MpqPath, archiveName)

		archive, err := am.loadArchive(archivePath)
		if err != nil {
			return err
		}

		am.entries = append(
			am.entries,
			archiveEntry{archivePath, archive.HashEntryMap},
		)
	}

	return nil
}
