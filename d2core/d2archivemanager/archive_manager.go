package d2archivemanager

import (
	"errors"
	"path"
	"sync"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2config"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2mpq"
)

type archiveEntry struct {
	archivePath  string
	hashEntryMap d2mpq.HashEntryMap
}

type ArchiveManager struct {
	cache   *d2common.Cache
	config  *d2config.Configuration
	entries []archiveEntry
	mutex   sync.Mutex
}

const (
	ArchiveBudget = 1024 * 1024 * 512
)

func CreateArchiveManager(config *d2config.Configuration) *ArchiveManager {
	return &ArchiveManager{cache: d2common.CreateCache(ArchiveBudget), config: config}
}

func (am *ArchiveManager) SetCacheVerbose(verbose bool) {
	am.cache.SetCacheVerbose(verbose)
}

func (am *ArchiveManager) GetCacheWeight() int {
	return am.cache.GetWeight()
}

func (am *ArchiveManager) GetCacheBudget() int {
	return am.cache.GetBudget()
}

func (am *ArchiveManager) ClearCache() {
	am.cache.Clear()
}

func (am *ArchiveManager) LoadArchiveForFile(filePath string) (*d2mpq.MPQ, error) {
	am.mutex.Lock()
	defer am.mutex.Unlock()

	if err := am.CacheArchiveEntries(); err != nil {
		return nil, err
	}

	for _, archiveEntry := range am.entries {
		if archiveEntry.hashEntryMap.Contains(filePath) {
			return am.LoadArchive(archiveEntry.archivePath)
		}
	}

	return nil, errors.New("file not found")
}

func (am *ArchiveManager) FileExistsInArchive(filePath string) (bool, error) {
	am.mutex.Lock()
	defer am.mutex.Unlock()

	if err := am.CacheArchiveEntries(); err != nil {
		return false, err
	}

	for _, archiveEntry := range am.entries {
		if archiveEntry.hashEntryMap.Contains(filePath) {
			return true, nil
		}
	}

	return false, nil
}

func (am *ArchiveManager) LoadArchive(archivePath string) (*d2mpq.MPQ, error) {
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

func (am *ArchiveManager) CacheArchiveEntries() error {
	if len(am.entries) == len(am.config.MpqLoadOrder) {
		return nil
	}

	am.entries = nil

	for _, archiveName := range am.config.MpqLoadOrder {
		archivePath := path.Join(am.config.MpqPath, archiveName)
		archive, err := am.LoadArchive(archivePath)
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
