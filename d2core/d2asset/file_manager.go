package d2asset

import (
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
)

const (
	fileBudget = 1024 * 1024 * 32
)

type fileManager struct {
	cache          *d2common.Cache
	archiveManager *archiveManager
	config         *d2config.Configuration
}

func createFileManager(config *d2config.Configuration, archiveManager *archiveManager) *fileManager {
	return &fileManager{d2common.CreateCache(fileBudget), archiveManager, config}
}

func (fm *fileManager) loadFile(filePath string) ([]byte, error) {
	filePath = fm.fixupFilePath(filePath)
	if value, found := fm.cache.Retrieve(filePath); found {
		return value.([]byte), nil
	}

	archive, err := fm.archiveManager.loadArchiveForFile(filePath)
	if err != nil {
		return nil, err
	}

	data, err := archive.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	if err := fm.cache.Insert(filePath, data, len(data)); err != nil {
		return nil, err
	}

	return data, nil
}

func (fm *fileManager) fileExists(filePath string) (bool, error) {
	filePath = fm.fixupFilePath(filePath)
	return fm.archiveManager.fileExistsInArchive(filePath)
}

func (fm *fileManager) fixupFilePath(filePath string) string {
	filePath = strings.ReplaceAll(filePath, "{LANG}", fm.config.Language)
	if strings.ToUpper(d2resource.LanguageCode) == "CHI" {
		filePath = strings.ReplaceAll(filePath, "{LANG_FONT}", fm.config.Language)
	} else {
		filePath = strings.ReplaceAll(filePath, "{LANG_FONT}", "latin")
	}

	filePath = strings.ToLower(filePath)
	filePath = strings.ReplaceAll(filePath, `/`, "\\")
	filePath = strings.TrimPrefix(filePath, "\\")

	return filePath
}
