package generator

import (
	"github.com/ThCompiler/go_game_constractor/scg/expr"
	"github.com/ThCompiler/go_game_constractor/scg/generator/codegen"
	servercodegen "github.com/ThCompiler/go_game_constractor/scg/server/codegen"
)

// Server add generate code of server
func Server(rootPkg, rootDir string, scriptInfo expr.ScriptInfo) ([]*codegen.File, error) {
	var files []*codegen.File

	files = append(files, servercodegen.ServerFile(rootPkg, rootDir, scriptInfo)...)
	files = append(files, servercodegen.ConfigFile(rootPkg, rootDir, scriptInfo)...)
	files = append(files, servercodegen.HandlerFile(rootPkg, rootDir, scriptInfo)...)
	files = append(files, servercodegen.AppFile(rootPkg, rootDir, scriptInfo)...)
	files = append(files, servercodegen.MainFile(rootPkg, rootDir, scriptInfo)...)
	files = append(files, servercodegen.LoggerPrepareFile(rootPkg, rootDir, scriptInfo)...)

	return files, nil
}
