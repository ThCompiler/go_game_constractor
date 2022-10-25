package generator

import (
	"github.com/thcompiler/go_game_constractor/scg/expr"
	"github.com/thcompiler/go_game_constractor/scg/generator/codegen"
	storecodegen "github.com/thcompiler/go_game_constractor/scg/store/codegen"
)

// Store add generate code of script store
func Store(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) ([]*codegen.File, error) {
	var files []*codegen.File

	files = append(files, storecodegen.ScriptStoreFiles(rootPkg, rootDir, scriptInfo)...)
	files = append(files, storecodegen.SaverStoreFile(rootPkg, rootDir, scriptInfo)...)
	files = append(files, storecodegen.PkgFiles(rootPkg, rootDir, scriptInfo)...)
	files = append(files, storecodegen.TextManagerFile(rootPkg, rootDir, scriptInfo)...)

	return files, nil
}
