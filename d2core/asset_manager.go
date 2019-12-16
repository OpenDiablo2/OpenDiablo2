package d2core

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/OpenDiablo2/D2Shared/d2common/d2resource"
	"github.com/OpenDiablo2/D2Shared/d2data/d2mpq"
	"github.com/OpenDiablo2/OpenDiablo2/d2corecommon"
)

type assetManager struct {
	fileCache     *cache
	archiveCache  *cache
	archiveLookup map[string]string
	config        *d2corecommon.Configuration
	mutex         sync.Mutex
}

func createAssetManager(config *d2corecommon.Configuration) *assetManager {
	return &assetManager{
		fileCache:     createCache(1024 * 1024 * 32),
		archiveCache:  createCache(1024 * 1024 * 512),
		archiveLookup: make(map[string]string),
		config:        config,
	}
}

func (am *assetManager) LoadFile(filePath string) []byte {
	data, err := am.loadFile(am.fixupFilePath(filePath))
	if err != nil {
		log.Println(err)
	}

	return data
}

func (am *assetManager) loadFile(filePath string) ([]byte, error) {
	if value, found := am.fileCache.retrieve(filePath); found {
		return value.([]byte), nil
	}

	archive, err := am.findArchiveForFilePath(filePath)
	if err != nil {
		return nil, err
	}

	data, err := archive.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	if err := am.fileCache.insert(filePath, data, len(data)); err != nil {
		return nil, err
	}

	return data, nil
}

func (am *assetManager) findArchiveForFilePath(filePath string) (*d2mpq.MPQ, error) {
	am.mutex.Lock()
	defer am.mutex.Unlock()

	if archivePath, found := am.archiveLookup[filePath]; found {
		return am.loadArchive(archivePath, true)
	}

	for _, archiveName := range am.config.MpqLoadOrder {
		archivePath := path.Join(am.config.MpqPath, archiveName)
		archive, err := am.loadArchive(archivePath, false)
		if err != nil {
			return nil, err
		}

		if archive.FileExists(filePath) {
			am.archiveLookup[filePath] = archivePath
			return archive, nil
		}
	}

	return nil, fmt.Errorf("file not found: %s", filePath)
}

func (am *assetManager) loadArchive(archivePath string, cache bool) (*d2mpq.MPQ, error) {
	if archive, found := am.archiveCache.retrieve(archivePath); found {
		return archive.(*d2mpq.MPQ), nil
	}

	archive, err := d2mpq.Load(archivePath)
	if err != nil {
		return nil, err
	}

	if cache {
		stat, err := os.Stat(archivePath)
		if err != nil {
			return nil, err
		}

		if err := am.archiveCache.insert(archivePath, archive, int(stat.Size())); err != nil {
			return nil, err
		}
	}

	return archive, nil
}

func (am *assetManager) fixupFilePath(filePath string) string {
	filePath = strings.ReplaceAll(filePath, "{LANG}", am.config.Language)
	if strings.ToUpper(d2resource.LanguageCode) == "CHI" {
		filePath = strings.ReplaceAll(filePath, "{LANG_FONT}", am.config.Language)
	} else {
		filePath = strings.ReplaceAll(filePath, "{LANG_FONT}", "latin")
	}

	filePath = strings.ToLower(filePath)
	filePath = strings.ReplaceAll(filePath, `/`, "\\")
	filePath = strings.TrimPrefix(filePath, "\\")

	return filePath
}
