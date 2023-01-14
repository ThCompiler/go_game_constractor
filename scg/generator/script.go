package generator

import (
	"github.com/ThCompiler/go_game_constractor/scg/expr"
	"github.com/ThCompiler/go_game_constractor/scg/generator/codegen"
	scriptcodegen "github.com/ThCompiler/go_game_constractor/scg/script/codegen"
)

// Script add generate code of script
func Script(rootPkg, rootDir string, scriptInfo expr.ScriptInfo) ([]*codegen.File, error) {
	var files []*codegen.File

	files = append(files, scriptcodegen.MatchersFile(rootPkg, rootDir, scriptInfo)...)
	files = append(files, scriptcodegen.ErrorsFile(rootPkg, rootDir, scriptInfo)...)
	files = append(files, scriptcodegen.PayloadsFile(rootPkg, rootDir, scriptInfo)...)
	files = append(files, scriptcodegen.ScriptFile(rootPkg, rootDir, scriptInfo)...)

	return files, nil
}
