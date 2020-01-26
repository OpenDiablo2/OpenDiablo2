package d2asset

import (
	"errors"
	"path"
	"sync"

	"github.com/OpenDiablo2/OpenDiablo2/d2corecommon"
	"github.com/OpenDiablo2/OpenDiablo2/d2data/d2mpq"
)

type archiveEntry struct {
	archivePath  string
	hashEntryMap d2mpq.HashEntryMap
}

type archiveManager struct {
	cache   *cache
	config  *d2corecommon.Configuration
	entries []archiveEntry
	mutex   sync.Mutex
}

func createArchiveManager(config *d2corecommon.Configuration) *archiveManager {
	return &archiveManager{cache: createCache(ArchiveBudget), config: config}
}

func (am *archiveManager) loadArchiveForFile(filePath string) (*d2mpq.MPQ, error) {
	am.mutex.Lock()
	defer am.mutex.Unlock()

	if err := am.cacheArchiveEntries(); err != nil {
		return nil, err
	}

	for _, archiveEntry := range am.entries {
		if archiveEntry.hashEntryMap.Contains(filePath) {
			return am.loadArchive(archiveEntry.archivePath)
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
	if archive, found := am.cache.retrieve(archivePath); found {
		return archive.(*d2mpq.MPQ), nil
	}

	archive, err := d2mpq.Load(archivePath)
	if err != nil {
		return nil, err
	}

	if err := am.cache.insert(archivePath, archive, int(archive.Data.ArchiveSize)); err != nil {
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
