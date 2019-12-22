package d2asset

import (
	"errors"

	"github.com/OpenDiablo2/D2Shared/d2data/d2cof"
	"github.com/OpenDiablo2/D2Shared/d2data/d2datadict"
	"github.com/OpenDiablo2/D2Shared/d2data/d2dc6"
	"github.com/OpenDiablo2/D2Shared/d2data/d2dcc"
	"github.com/OpenDiablo2/D2Shared/d2data/d2mpq"
	"github.com/OpenDiablo2/OpenDiablo2/d2corecommon"
)

const (
	// In megabytes
	ArchiveBudget = 1024 * 1024 * 512
	FileBudget    = 1024 * 1024 * 32

	// In counts
	PaletteBudget   = 64
	PaperdollBudget = 64
	AnimationBudget = 64
)

var (
	ErrHasInit error = errors.New("asset system is already initialized")
	ErrNoInit  error = errors.New("asset system is not initialized")
)

type assetManager struct {
	archiveManager   *archiveManager
	fileManager      *fileManager
	paletteManager   *paletteManager
	paperdollManager *paperdollManager
	animationManager *animationManager
}

var singleton *assetManager

func Initialize(config *d2corecommon.Configuration) error {
	if singleton != nil {
		return ErrHasInit
	}

	var (
		archiveManager   = createArchiveManager(config)
		fileManager      = createFileManager(config, archiveManager)
		paletteManager   = createPaletteManager()
		paperdollManager = createPaperdollManager()
		animationManager = createAnimationManager()
	)

	singleton = &assetManager{
		archiveManager,
		fileManager,
		paletteManager,
		paperdollManager,
		animationManager,
	}

	return nil
}

func LoadArchive(archivePath string) (*d2mpq.MPQ, error) {
	if singleton == nil {
		return nil, ErrNoInit
	}

	return singleton.archiveManager.loadArchive(archivePath)
}

func LoadFile(filePath string) ([]byte, error) {
	if singleton == nil {
		return nil, ErrNoInit
	}

	return singleton.fileManager.loadFile(filePath)
}

func LoadAnimation(animationPath, palettePath string) (*Animation, error) {
	if singleton == nil {
		return nil, ErrNoInit
	}

	return singleton.animationManager.loadAnimation(animationPath, palettePath)
}

func LoadPaperdoll(object *d2datadict.ObjectLookupRecord, palettePath string) (*Paperdoll, error) {
	if singleton == nil {
		return nil, ErrNoInit
	}

	return singleton.paperdollManager.loadPaperdoll(object, palettePath)
}

// TODO: remove transitional usage pattern
func MustLoadFile(filePath string) []byte {
	data, err := LoadFile(filePath)
	if err != nil {
		return []byte{}
	}

	return data
}

func loadPalette(palettePath string) (*d2datadict.PaletteRec, error) {
	if singleton == nil {
		return nil, ErrNoInit
	}

	return singleton.paletteManager.loadPalette(palettePath)
}

func loadDC6(dc6Path, palettePath string) (*d2dc6.DC6File, error) {
	dc6Data, err := LoadFile(dc6Path)
	if err != nil {
		return nil, err
	}

	paletteData, err := loadPalette(palettePath)
	if err != nil {
		return nil, err
	}

	dc6, err := d2dc6.LoadDC6(dc6Data, *paletteData)
	if err != nil {
		return nil, err
	}

	return &dc6, nil
}

func LoadDCC(dccPath string) (*d2dcc.DCC, error) {
	dccData, err := LoadFile(dccPath)
	if err != nil {
		return nil, err
	}

	return d2dcc.LoadDCC(dccData)
}

func LoadCOF(cofPath string) (*d2cof.COF, error) {
	cofData, err := LoadFile(cofPath)
	if err != nil {
		return nil, err
	}

	return d2cof.LoadCOF(cofData)
}
