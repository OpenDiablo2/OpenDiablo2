package d2systems

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2mpq"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

const (
	logPrefixFileSourceResolver = "ComponentFactory Source Resolver"
)

// FileSourceResolver is responsible for determining if files can be used as a file source.
// If it can, it sets the file up as a source, and the file handle resolver system can
// then use the source to open files.
type FileSourceResolver struct {
	akara.BaseSubscriberSystem
	*d2util.Logger
	filesToCheck   *akara.Subscription
	Components struct {
		File d2components.FileFactory
		FileType d2components.FileTypeFactory
		FileSource d2components.FileSourceFactory
	}
}

// Init initializes the file source resolver, injecting the necessary components into the world
func (m *FileSourceResolver) Init(world *akara.World) {
	m.World = world

	m.setupLogger()

	m.Debug("initializing ...")

	m.setupSubscriptions()
	m.setupFactories()

	m.Debug("... initialization complete!")
}

func (m *FileSourceResolver) setupLogger() {
	m.Logger = d2util.NewLogger()
	m.SetPrefix(logPrefixFileSourceResolver)
}

func (m *FileSourceResolver) setupSubscriptions() {
	m.Debug("setting up component subscriptions")

	// subscribe to entities with a file type and file path, but no file source type
	filesToCheck := m.NewComponentFilter().
		Require(
			&d2components.File{},
			&d2components.FileType{},
		).
		Forbid(
			&d2components.FileSource{},
		).
		Build()

	m.filesToCheck = m.World.AddSubscription(filesToCheck)
}

func (m *FileSourceResolver) setupFactories() {
	m.Debug("setting up component factories")

	m.InjectComponent(&d2components.File{}, &m.Components.File.ComponentFactory)
	m.InjectComponent(&d2components.FileType{}, &m.Components.FileType.ComponentFactory)
	m.InjectComponent(&d2components.FileSource{}, &m.Components.FileSource.ComponentFactory)
}

// Update iterates over entities from its subscription, and checks if it can be used as a file source
func (m *FileSourceResolver) Update() {
	for _, eid := range m.filesToCheck.GetEntities() {
		m.processSourceEntity(eid)
	}
}

func (m *FileSourceResolver) processSourceEntity(id akara.EID) {
	fp, found := m.Components.File.Get(id)
	if !found {
		return
	}

	ft, found := m.Components.FileType.Get(id)
	if !found {
		return
	}

	switch ft.Type {
	case d2enum.FileTypeUnknown:
		m.Errorf("unknown file type for file `%s`", fp.Path)
		return
	case d2enum.FileTypeMPQ:
		instance, err := m.makeMpqSource(fp.Path)

		if err != nil {
			ft.Type = d2enum.FileTypeUnknown
			break
		}

		m.Components.FileSource.Add(id).AbstractSource = instance

		m.Debugf("adding MPQ source: `%s`", fp.Path)
	case d2enum.FileTypeDirectory:
		m.Components.FileSource.Add(id).AbstractSource = m.makeFileSystemSource(fp.Path)
		m.Debugf("adding FILESYSTEM source: `%s`", fp.Path)
	}
}

// filesystem source
func (m *FileSourceResolver) makeFileSystemSource(path string) d2components.AbstractSource {
	return &fsSource{rootDir: path}
}

type fsSource struct {
	rootDir string
}

func (s *fsSource) Open(path *d2components.File) (d2interface.DataStream, error) {
	fileData, err := os.Open(s.fullPath(path.Path))
	if err != nil {
		return nil, err
	}

	return fileData, nil
}

func (s *fsSource) fullPath(path string) string {
	return filepath.Clean(filepath.Join(s.rootDir, path))
}

func (s *fsSource) Path() string {
	return filepath.Clean(s.rootDir)
}

// mpq source
func (m *FileSourceResolver) makeMpqSource(path string) (d2components.AbstractSource, error) {
	mpq, err := d2mpq.Load(path)
	if err != nil {
		return nil, err
	}

	return &mpqSource{mpq: mpq}, nil
}

type mpqSource struct {
	mpq d2interface.Archive
}

func (s *mpqSource) Open(path *d2components.File) (d2interface.DataStream, error) {
	fileData, err := s.mpq.ReadFileStream(s.cleanMpqPath(path.Path))
	if err != nil {
		return nil, err
	}

	return fileData, nil
}

func (s *mpqSource) cleanMpqPath(path string) string {
	path = strings.ReplaceAll(path, "/", "\\")

	if string(path[0]) == "\\" {
		path = path[1:]
	}

	return path
}

func (s *mpqSource) Path() string {
	return filepath.Clean(s.mpq.Path())
}
