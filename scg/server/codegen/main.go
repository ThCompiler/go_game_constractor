package codegen

import (
	"path"
	"path/filepath"
	"strings"

	"github.com/ThCompiler/go_game_constractor/scg/expr"
	"github.com/ThCompiler/go_game_constractor/scg/generator/codegen"
)

// MainFile returns main file
func MainFile(rootPkg, rootDir string, scriptInfo expr.ScriptInfo) []*codegen.File {
	mainFile := generateMain(rootPkg, rootDir, scriptInfo)

	return []*codegen.File{mainFile}
}

func generateMain(rootPkg, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
	var sections []*codegen.SectionTemplate

	fpath := filepath.Join(rootDir, "cmd", "main.go")
	imports := []*codegen.ImportSpec{
		{Path: path.Join("log")},
		codegen.SCGImport(path.Join("pkg", "convertor", "words")),
		codegen.SCGImport(path.Join("pkg", "convertor", "words", "languages")),
		{Path: path.Join(rootPkg, "config")},
		{Path: path.Join(rootPkg, "internal", strings.ToLower(codegen.ToTitle(scriptInfo.Name)))},
	}

	sections = []*codegen.SectionTemplate{
		codegen.Header(codegen.ToTitle(scriptInfo.Name)+"-Main file", "main", imports, true),
	}

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "main-file",
		Source: mainFileStructT,
		Data:   scriptInfo,
		FuncMap: map[string]interface{}{
			"ToTitle": codegen.ToTitle,
			"ToLower": strings.ToLower,
		},
	})

	return &codegen.File{Path: fpath, SectionTemplates: sections, IsUpdatable: true}
}

const mainFileStructT = ` // This is your application's startup file. 
// All the basic settings take place already in the "internal/app" file. 

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	_ = words.LoadWordsConstants(languages.Russia, cfg.App.ResourcesDir)

	// Run
	{{ ToLower ( ToTitle .Name ) }}.Run(cfg)
}
`
