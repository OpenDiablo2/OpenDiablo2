package d2filemanager

import (
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2archivemanager"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2config"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
)

const (
	FileBudget = 1024 * 1024 * 32
)

type FileManager struct {
	cache          *d2common.Cache
	archiveManager *d2archivemanager.ArchiveManager
	config         *d2config.Configuration
}

func CreateFileManager(config *d2config.Configuration, archiveManager *d2archivemanager.ArchiveManager) *FileManager {
	return &FileManager{d2common.CreateCache(FileBudget), archiveManager, config}
}

func (fm *FileManager) SetCacheVerbose(verbose bool) {
	fm.cache.SetCacheVerbose(verbose)
}

func (fm *FileManager) ClearCache() {
	fm.cache.Clear()
}

func (fm *FileManager) GetCacheWeight() int {
	return fm.cache.GetWeight()
}

func (fm *FileManager) GetCacheBudget() int {
	return fm.cache.GetBudget()
}
func (fm *FileManager) LoadFile(filePath string) ([]byte, error) {
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

func (fm *FileManager) FileExists(filePath string) (bool, error) {
	filePath = fm.fixupFilePath(filePath)
	return fm.archiveManager.FileExistsInArchive(filePath)
}

func (fm *FileManager) fixupFilePath(filePath string) string {
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
