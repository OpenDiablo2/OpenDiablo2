package d2asset

import (
	"errors"
	"path"
	"sync"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2mpq"
)

type archiveManager struct {
	cache    d2interface.Cache
	config   d2interface.Configuration
	archives []d2interface.Archive
	mutex    sync.Mutex
}

const (
	archiveBudget = 1024 * 1024 * 512
)

func createArchiveManager(config d2interface.Configuration) *archiveManager {
	return &archiveManager{cache: d2common.CreateCache(archiveBudget), config: config}
}

func (am *archiveManager) loadArchiveForFile(filePath string) (d2interface.Archive, error) {
	am.mutex.Lock()
	defer am.mutex.Unlock()

	if err := am.cacheArchiveEntries(); err != nil {
		return nil, err
	}

	for _, archive := range am.archives {
		if archive.Contains(filePath) {
			result, ok := am.loadArchive(archive.Path())
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

	for _, archiveEntry := range am.archives {
		if archiveEntry.Contains(filePath) {
			return true, nil
		}
	}

	return false, nil
}

func (am *archiveManager) loadArchive(archivePath string) (d2interface.Archive, error) {
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

func (am *archiveManager) cacheArchiveEntries() error {
	if len(am.archives) == len(am.config.MpqLoadOrder()) {
		return nil
	}

	am.archives = nil

	for _, archiveName := range am.config.MpqLoadOrder() {
		archivePath := path.Join(am.config.MpqPath(), archiveName)

		archive, err := am.loadArchive(archivePath)
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
