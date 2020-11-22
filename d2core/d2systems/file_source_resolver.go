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

// NewFileSourceResolver creates a new file source resolver system
func NewFileSourceResolver() *FileSourceResolver {
	// subscribe to entities with a file type and file path, but no file source type
	filesToCheck := akara.NewFilter().
		Require(d2components.FilePath).
		Require(d2components.FileType).
		Forbid(d2components.FileSource).
		Build()

	fsr := &FileSourceResolver{
		BaseSubscriberSystem: akara.NewBaseSubscriberSystem(filesToCheck),
		Logger:               d2util.NewLogger(),
	}

	fsr.SetPrefix(logPrefixFileSourceResolver)

	return fsr
}

// FileSourceResolver is responsible for determining if files can be used as a file source.
// If it can, it sets the file up as a source, and the file handle resolver system can
// then use the source to open files.
type FileSourceResolver struct {
	*akara.BaseSubscriberSystem
	*d2util.Logger
	fileSub     *akara.Subscription
	filePaths   *d2components.FilePathMap
	fileTypes   *d2components.FileTypeMap
	fileSources *d2components.FileSourceMap
}

// Init initializes the file source resolver, injecting the necessary components into the world
func (m *FileSourceResolver) Init(_ *akara.World) {
	m.Info("initializing ...")

	m.fileSub = m.Subscriptions[0]

	// try to inject the components we require, then cast the returned
	// abstract ComponentMap back to the concrete implementation
	m.filePaths = m.InjectMap(d2components.FilePath).(*d2components.FilePathMap)
	m.fileTypes = m.InjectMap(d2components.FileType).(*d2components.FileTypeMap)
	m.fileSources = m.InjectMap(d2components.FileSource).(*d2components.FileSourceMap)
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
	fp, found := m.filePaths.GetFilePath(id)
	if !found {
		return
	}

	ft, found := m.fileTypes.GetFileType(id)
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

		source := m.fileSources.AddFileSource(id)
		source.AbstractSource = instance

		m.Infof("using MPQ source for `%s`", fp.Path)
	case d2enum.FileTypeDirectory:
		instance := m.makeFileSystemSource(fp.Path)

		source := m.fileSources.AddFileSource(id)
		source.AbstractSource = instance

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

func (s *fsSource) Open(path *d2components.FilePathComponent) (d2interface.DataStream, error) {
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

func (s *mpqSource) Open(path *d2components.FilePathComponent) (d2interface.DataStream, error) {
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
