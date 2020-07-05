package d2asset

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
)

const (
	fileBudget = 1024 * 1024 * 32
)

type fileManager struct {
	cache          d2interface.Cache
	archiveManager d2interface.ArchiveManager
	config         d2interface.Configuration
}

func createFileManager(config d2interface.Configuration,
	archiveManager d2interface.ArchiveManager) *fileManager {
	return &fileManager{
		d2common.CreateCache(fileBudget),
		archiveManager,
		config,
	}
}

func (fm *fileManager) loadFileStream(filePath string) (d2interface.ArchiveDataStream, error) {
	filePath = fm.fixupFilePath(filePath)

	archive, err := fm.archiveManager.LoadArchiveForFile(filePath)
	if err != nil {
		return nil, err
	}

	return archive.ReadFileStream(filePath)
}

func (fm *fileManager) loadFile(filePath string) ([]byte, error) {
	filePath = fm.fixupFilePath(filePath)
	if value, found := fm.cache.Retrieve(filePath); found {
		return value.([]byte), nil
	}

	archive, err := fm.archiveManager.LoadArchiveForFile(filePath)
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
	return fm.archiveManager.FileExistsInArchive(filePath)
}

func (fm *fileManager) fixupFilePath(filePath string) string {
	filePath = fm.removeLocaleTokens(filePath)
	filePath = strings.ToLower(filePath)
	filePath = strings.ReplaceAll(filePath, `/`, "\\")
	filePath = strings.TrimPrefix(filePath, "\\")

	return filePath
}

func (fm *fileManager) removeLocaleTokens(filePath string) string {
	tableToken := d2resource.LanguageTableToken
	fontToken := d2resource.LanguageFontToken

	filePath = strings.ReplaceAll(filePath, tableToken, fm.config.Language())

	// fixme: not all languages==latin
	filePath = strings.ReplaceAll(filePath, fontToken, "latin")

	return filePath
}
