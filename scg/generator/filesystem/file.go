package filesystem

import (
    "bytes"
    "github.com/c2fo/vfs/v6"
    "github.com/google/uuid"
    "github.com/pkg/errors"
    "os"
    "path/filepath"
    "unsafe"
)

// copy from system lib for construct system function on vfs.File

// Slice is the runtime representation of a slice.
// It cannot be used safely or portably and its representation may
// change in a later release.
//
// Unlike reflect.SliceHeader, its Data field is sufficient to guarantee the
// data it references will not be garbage collected.
type Slice struct {
    Data unsafe.Pointer
    Len  int
    Cap  int
}

// String is the runtime representation of a string.
// It cannot be used safely or portably and its representation may
// change in a later release.
//
// Unlike reflect.StringHeader, its Data field is sufficient to guarantee the
// data it references will not be garbage collected.
type String struct {
    Data unsafe.Pointer
    Len  int
}

func WriteString(file vfs.File, s string) (n int, err error) {
    var b []byte
    hdr := (*Slice)(unsafe.Pointer(&b))
    hdr.Data = (*String)(unsafe.Pointer(&s)).Data
    hdr.Cap = len(s)
    hdr.Len = len(s)
    return file.Write(b)
}

func MkDirAll(loc vfs.Location, path string) (vfs.Location, error) {
    return loc.NewLocation(path)
}

func ReadFile(loc vfs.Location, path string) (*bytes.Buffer, error) {
    src, err := loc.NewFile(path)
    if err != nil {
        return nil, err
    }

    buf, err := ReadToBuffer(src)
    if err != nil {
        return nil, err
    }

    if err1 := src.Close(); err1 != nil && err == nil {
        err = err1
    }

    return buf, nil
}

func WriteFile(loc vfs.Location, name string, data []byte) error {
    src, err := loc.NewFile(name)
    if err != nil {
        return err
    }

    _, err = src.Write(data)

    if err1 := src.Close(); err1 != nil && err == nil {
        err = err1
    }
    return err
}

func ReadToBuffer(file vfs.File) (*bytes.Buffer, error) {
    buf := bytes.NewBuffer(make([]byte, 0))
    _, err := buf.ReadFrom(file)
    if err != nil {
        return nil, err
    }
    return buf, nil
}

type TempFile struct {
    vfs.File
    name string
}

func (tf *TempFile) Name() string {
    return tf.name
}

func NewTempFile(loc vfs.Location, dir string, pattern string) (*TempFile, error) {
    if dir == "" {
        dir = os.TempDir()
    }

    prefix, suffix, err := prefixAndSuffix(pattern)
    if err != nil {
        return nil, &os.PathError{Op: "createtemp", Path: pattern, Err: err}
    }
    prefix = filepath.Clean(prefix)

    tempLoc, err := loc.NewLocation(dir)
    if err != nil {
        return nil, &os.PathError{Op: "createtemp", Path: pattern, Err: err}
    }

    try := 0
    for {
        name := prefix + nextRandom() + suffix
        f, err := tempLoc.NewFile(name)
        if err != nil {
            return nil, &os.PathError{Op: "createtemp", Path: pattern, Err: err}
        }

        is, err := f.Exists()
        if err != nil {
            return nil, &os.PathError{Op: "createtemp", Path: pattern, Err: err}
        }

        if is {
            if try++; try < 10000 {
                continue
            }
            return nil, &os.PathError{Op: "createtemp", Path: prefix + "*" + suffix, Err: os.ErrExist}
        }

        return &TempFile{
            File: f,
            name: name,
        }, err
    }
}

func nextRandom() string {
    return uuid.New().String()
}

var errPatternHasSeparator = errors.New("pattern contains path separator")

func prefixAndSuffix(pattern string) (prefix, suffix string, err error) {
    for i := 0; i < len(pattern); i++ {
        if os.IsPathSeparator(pattern[i]) {
            return "", "", errPatternHasSeparator
        }
    }
    if pos := lastIndex(pattern, '*'); pos != -1 {
        prefix, suffix = pattern[:pos], pattern[pos+1:]
    } else {
        prefix = pattern
    }
    return prefix, suffix, nil
}

func lastIndex(s string, sep byte) int {
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == sep {
            return i
        }
    }
    return -1
}
