package d2systems

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"os"
	"path/filepath"
	"strings"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2mpq"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

const (
	logPrefixFileTypeResolver = "File Type Resolver"
)

// NewFileTypeResolver creates a new file type resolution system.
func NewFileTypeResolver() *FileTypeResolver {
	// we subscribe only to entities that have a filepath
	// and have not yet been given a file type
	filesToCheck := akara.NewFilter().
		Require(d2components.FilePath).
		Forbid(d2components.FileType).
		Build()

	ftr := &FileTypeResolver{
		SubscriberSystem: akara.NewSubscriberSystem(filesToCheck),
		Logger: d2util.NewLogger(),
	}

	ftr.SetPrefix(logPrefixFileTypeResolver)

	return ftr
}

// static check that FileTypeResolver implements the System interface
var _ akara.System = &FileTypeResolver{}

// FileTypeResolver is responsible for determining file types from file paths.
// This system will subscribe to entities that have a file path component, but do not
// have a file type component. It will use the file path component to determine the file type,
// and it will then create the file type component for the entity, thus removing the entity
// from its subscription.
type FileTypeResolver struct {
	*akara.SubscriberSystem
	*d2util.Logger
	filesToCheck *akara.Subscription
	filePaths    *d2components.FilePathMap
	fileTypes    *d2components.FileTypeMap
}

// Init initializes the system with the given world
func (m *FileTypeResolver) Init(world *akara.World) {
	m.World = world

	if world == nil {
		m.SetActive(false)
		return
	}

	m.Info("initializing ...")

	for subIdx := range m.Subscriptions {
		m.Subscriptions[subIdx] = m.AddSubscription(m.Subscriptions[subIdx].Filter)
	}

	m.filesToCheck = m.Subscriptions[0]

	// try to inject the components we require, then cast the returned
	// abstract ComponentMap back to the concrete implementation
	m.filePaths = m.InjectMap(d2components.FilePath).(*d2components.FilePathMap)
	m.fileTypes = m.InjectMap(d2components.FileType).(*d2components.FileTypeMap)
}

// Process processes all of the Entities
func (m *FileTypeResolver) Process() {
	for _, eid := range m.filesToCheck.GetEntities() {
		m.determineFileType(eid)
	}
}

func (m *FileTypeResolver) determineFileType(id akara.EID) {
	fp, found := m.filePaths.GetFilePath(id)
	if !found {
		return
	}

	ft := m.fileTypes.AddFileType(id)
	if _, err := d2mpq.Load(fp.Path); err == nil {
		ft.Type = d2enum.FileTypeMPQ
		return
	}

	ext := strings.ToLower(filepath.Ext(fp.Path))

	switch ext {
	case ".mpq":
		ft.Type = d2enum.FileTypeMPQ
	case ".d2":
		ft.Type = d2enum.FileTypeD2
	case ".dcc":
		ft.Type = d2enum.FileTypeDCC
	case ".dc6":
		ft.Type = d2enum.FileTypeDC6
	case ".wav":
		ft.Type = d2enum.FileTypeWAV
	case ".ds1":
		ft.Type = d2enum.FileTypeDS1
	case ".dt1":
		ft.Type = d2enum.FileTypeDT1
	case ".pl2":
		ft.Type = d2enum.FileTypePaletteTransform
	case ".dat":
		ft.Type = d2enum.FileTypePalette
	case ".tbl":
		ft.Type = d2enum.FileTypeStringTable
		// HACK: we should probably not use the path to check for the type
		// but we have two types of .tbl file :(
		if strings.Contains(fp.Path, "FONT") {
			ft.Type = d2enum.FileTypeFontTable
		}
	case ".txt":
		ft.Type = d2enum.FileTypeDataDictionary
	case ".cof":
		ft.Type = d2enum.FileTypeCOF
	case ".json":
		ft.Type = d2enum.FileTypeJSON
	default:
		cleanPath := filepath.Clean(fp.Path)

		info, err := os.Lstat(cleanPath)
		if err != nil {
			ft.Type = d2enum.FileTypeUnknown
			return
		}

		if info.Mode().IsDir() {
			ft.Type = d2enum.FileTypeDirectory
			return
		}
	}
}
