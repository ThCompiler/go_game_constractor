package filesystem

import (
    "github.com/c2fo/vfs/v6"
    "path/filepath"
)

type Filesystem struct {
    fileSystem vfs.FileSystem
    volume     string
}

func NewFilesystem(fs vfs.FileSystem, volume string) *Filesystem {
    return &Filesystem{
        fileSystem: fs,
        volume:     volume,
    }
}

// NewFile initializes a File on the specified volume at path 'absFilePath'.
//
//   * Accepts volume and an absolute file path.
//   * Upon success, a vfs.File, representing the file's new path (location path + file relative path), will be returned.
//   * On error, nil is returned for the file.
//   * Note that not all file systems will have a "volume" and will therefore be "":
//       file:///path/to/file has a volume of "" and name /path/to/file
//     whereas
//       s3://mybucket/path/to/file has a volume of "mybucket and name /path/to/file
//     results in /tmp/dir1/newerdir/file.txt for the final vfs.File path.
//   * The file may or may not already exist.
func (vf *Filesystem) NewFile(absFilePath string) (vfs.File, error) {
    return vf.fileSystem.NewFile(vf.volume, filepath.FromSlash(filepath.Clean("/"+absFilePath)))
}

// NewLocation initializes a Location on the specified volume with the given path.
//
//   * Accepts volume and an absolute location path.
//   * The file may or may not already exist. Note that on key-store file systems like S3 or GCS, paths never truly exist.
//   * On error, nil is returned for the location.
//
// See NewFile for note on volume.
func (vf *Filesystem) NewLocation(absLocPath string) (vfs.Location, error) {
    loc, err := vf.fileSystem.NewLocation(vf.volume, filepath.FromSlash(filepath.Clean("/"+absLocPath)+"/"))
    return &Location{loc}, err
}

// Name returns the name of the FileSystem ie: Amazon S3, os, Google Cloud Storage, etc.
func (vf *Filesystem) Name() string {
    return vf.fileSystem.Name()
}

// Scheme returns the uri scheme used by the FileSystem: s3, file, gs, etc.
func (vf *Filesystem) Scheme() string {
    return vf.fileSystem.Scheme()
}

// Retry will return the retry function to be used by any file system.
func (vf *Filesystem) Retry() vfs.Retry {
    return vf.fileSystem.Retry()
}

type Location struct {
    vfs.Location
}

// NewLocation add '/' to end of path for run fs.Location.NewLocation.
func (l *Location) NewLocation(relLocPath string) (vfs.Location, error) {
    loc, err := l.Location.NewLocation(filepath.FromSlash(filepath.Clean(relLocPath) + "/"))
    return &Location{loc}, err
}

// ChangeDir add '/' to end of path for run fs.Location.NewLocation.
func (l *Location) ChangeDir(relLocPath string) error {
    err := l.Location.ChangeDir(filepath.FromSlash(filepath.Clean(relLocPath) + "/"))
    return err
}
