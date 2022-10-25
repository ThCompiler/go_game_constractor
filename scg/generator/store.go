package generator

import (
	"gameconstractor/scg/expr"
	"gameconstractor/scg/generator/codegen"
	storecodegen "gameconstractor/scg/store/codegen"
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
