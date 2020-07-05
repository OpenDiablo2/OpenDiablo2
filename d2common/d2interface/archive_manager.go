package d2interface

// ArchiveManager manages loading files from archives
type ArchiveManager interface {
	LoadArchiveForFile(filePath string) (Archive, error)
	FileExistsInArchive(filePath string) (bool, error)
	LoadArchive(archivePath string) (Archive, error)
	CacheArchiveEntries() error
	SetVerbose(verbose bool)
	ClearCache()
	GetCache() Cache
}
