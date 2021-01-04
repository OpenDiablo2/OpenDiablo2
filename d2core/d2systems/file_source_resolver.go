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
	logPrefixFileSourceResolver = "File Source Resolver"
)

// FileSourceResolver is responsible for determining if files can be used as a file source.
// If it can, it sets the file up as a source, and the file handle resolver system can
// then use the source to open files.
type FileSourceResolver struct {
	akara.BaseSubscriberSystem
	*d2util.Logger
	filesToCheck *akara.Subscription
	d2components.FilePathFactory
	d2components.FileTypeFactory
	d2components.FileSourceFactory
}

// Init initializes the file source resolver, injecting the necessary components into the world
func (m *FileSourceResolver) Init(world *akara.World) {
	m.World = world

	m.setupLogger()

	m.Info("initializing ...")

	m.setupSubscriptions()
	m.setupFactories()

	m.Info("... initialization complete!")
}

func (m *FileSourceResolver) setupLogger() {
	m.Logger = d2util.NewLogger()
	m.SetPrefix(logPrefixFileSourceResolver)
}

func (m *FileSourceResolver) setupSubscriptions() {
	m.Info("setting up component subscriptions")

	// subscribe to entities with a file type and file path, but no file source type
	filesToCheck := m.NewComponentFilter().
		Require(
			&d2components.FilePath{},
			&d2components.FileType{},
		).
		Forbid(
			&d2components.FileSource{},
		).
		Build()

	m.filesToCheck = m.AddSubscription(filesToCheck)
}

func (m *FileSourceResolver) setupFactories() {
	m.Info("setting up component factories")

	filePathID := m.RegisterComponent(&d2components.FilePath{})
	fileTypeID := m.RegisterComponent(&d2components.FileType{})
	fileSourceID := m.RegisterComponent(&d2components.FileSource{})

	m.FilePath = m.GetComponentFactory(filePathID)
	m.FileType = m.GetComponentFactory(fileTypeID)
	m.FileSource = m.GetComponentFactory(fileSourceID)
}

// Update iterates over entities from its subscription, and checks if it can be used as a file source
func (m *FileSourceResolver) Update() {
	for subIdx := range m.Subscriptions {
		for _, sourceEntityID := range m.Subscriptions[subIdx].GetEntities() {
			m.processSourceEntity(sourceEntityID)
		}
	}
}

func (m *FileSourceResolver) processSourceEntity(id akara.EID) {
	fp, found := m.GetFilePath(id)
	if !found {
		return
	}

	ft, found := m.GetFileType(id)
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

		m.AddFileSource(id).AbstractSource = instance

		m.Infof("using MPQ source for `%s`", fp.Path)
	case d2enum.FileTypeDirectory:
		m.AddFileSource(id).AbstractSource = m.makeFileSystemSource(fp.Path)
		m.Infof("using FILESYSTEM source for `%s`", fp.Path)
	}
}

// filesystem source
func (m *FileSourceResolver) makeFileSystemSource(path string) d2components.AbstractSource {
	return &fsSource{rootDir: path}
}

type fsSource struct {
	rootDir string
}

func (s *fsSource) Open(path *d2components.FilePath) (d2interface.DataStream, error) {
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

func (s *mpqSource) Open(path *d2components.FilePath) (d2interface.DataStream, error) {
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
