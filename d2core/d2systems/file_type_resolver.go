package d2systems

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"os"
	"path/filepath"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2mpq"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

const (
	logPrefixFileTypeResolver = "ComponentFactory Type Resolver"
)

// static check that FileTypeResolver implements the System interface
var _ akara.System = &FileTypeResolver{}

// FileTypeResolver is responsible for determining file types from file paths.
// This system will subscribe to entities that have a file path component, but do not
// have a file type component. It will use the file path component to determine the file type,
// and it will then create the file type component for the entity, thus removing the entity
// from its subscription.
type FileTypeResolver struct {
	akara.BaseSubscriberSystem
	*d2util.Logger
	filesToCheck *akara.Subscription
	Components struct {
		File d2components.FileFactory
		FileType d2components.FileTypeFactory
	}
}

// Init initializes the system with the given world
func (m *FileTypeResolver) Init(world *akara.World) {
	m.World = world

	m.setupLogger()

	m.Debug("initializing ...")

	m.setupFactories()
	m.setupSubscriptions()
}

func (m *FileTypeResolver) setupLogger() {
	m.Logger = d2util.NewLogger()
	m.SetPrefix(logPrefixFileTypeResolver)
}

func (m *FileTypeResolver) setupFactories() {
	m.InjectComponent(&d2components.File{}, &m.Components.File.ComponentFactory)
	m.InjectComponent(&d2components.FileType{}, &m.Components.FileType.ComponentFactory)
}

func (m *FileTypeResolver) setupSubscriptions() {
	// we subscribe only to entities that have a filepath
	// and have not yet been given a file type
	filesToCheck := m.NewComponentFilter().
		Require(&d2components.File{}).
		Forbid(&d2components.FileType{}).
		Build()

	m.filesToCheck = m.AddSubscription(filesToCheck)
}

// Update processes all of the Entities
func (m *FileTypeResolver) Update() {
	for _, eid := range m.filesToCheck.GetEntities() {
		m.determineFileType(eid)
	}
}

//nolint:gocyclo // this big switch statement is unfortunate, but necessary
func (m *FileTypeResolver) determineFileType(id akara.EID) {
	fp, found := m.Components.File.Get(id)
	if !found {
		return
	}

	ft := m.Components.FileType.Add(id)

	// try to immediately load as an mpq
	if _, err := d2mpq.Load(fp.Path); err == nil {
		ft.Type = d2enum.FileTypeMPQ
		return
	}

	// special case for the locale file
	if fp.Path == d2resource.LocalLanguage {
		ft.Type = d2enum.FileTypeLocale
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
