package generator

import (
	"github.com/thcompiler/go_game_constractor/scg/expr"
	"github.com/thcompiler/go_game_constractor/scg/generator/codegen"
)

// Genfunc is the type of the functions invoked to generate code.
type Genfunc func(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) ([]*codegen.File, error)

// Generators returns the qualified paths (including the package name) to the
// code generator functions for the given command, an error if the command is
// not supported. Generators is a public variable so that external code (e.g.
// plugins) may override the default generators.
var Generators = generators

// generators returns the generator functions exposed by the generator
func generators() []Genfunc {
	return []Genfunc{Store}
}
