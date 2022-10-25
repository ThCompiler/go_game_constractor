package codegen

import (
	"fmt"
)

type (
	// ImportSpec defines a generated import statement.
	ImportSpec struct {
		// Name of imported package if needed.
		Name string
		// Go import path of package.
		Path string
	}
)

// NewImport creates an import spec.
func NewImport(name, path string) *ImportSpec {
	return &ImportSpec{Name: name, Path: path}
}

// SimpleImport creates an import with no explicit path component.
func SimpleImport(path string) *ImportSpec {
	return &ImportSpec{Path: path}
}

// SCGImport creates an import for a Goa package.
func SCGImport(rel string) *ImportSpec {
	name := ""
	if rel == "" {
		name = "scg"
		rel = "pkg"
	}
	return SCGNamedImport(rel, name)
}

// SCGNamedImport creates an import for a Goa package with the given name.
func SCGNamedImport(rel, name string) *ImportSpec {
	root := "github.com/ThCompiler/go_game_constractor"
	if rel != "" {
		rel = "/" + rel
	}
	return &ImportSpec{Name: name, Path: root + rel}
}

// Code returns the Go import statement for the ImportSpec.
func (s *ImportSpec) Code() string {
	if len(s.Name) > 0 {
		return fmt.Sprintf(`%s "%s"`, s.Name, s.Path)
	}
	return fmt.Sprintf(`"%s"`, s.Path)
}
