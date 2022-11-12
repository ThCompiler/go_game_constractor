package codegen

import (
	"github.com/ThCompiler/go_game_constractor/scg/expr"
	"github.com/ThCompiler/go_game_constractor/scg/generator/codegen"
	"path"
	"path/filepath"
)

// MainFile returns main file
func MainFile(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) []*codegen.File {
	mainFile := generateMain(rootPkg, rootDir, scriptInfo)

	return []*codegen.File{mainFile}
}

func generateMain(_ string, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
	var sections []*codegen.SectionTemplate

	fpath := filepath.Join(rootDir, "cmd", "main.go")
	imports := []*codegen.ImportSpec{
		codegen.SCGImport(path.Join("director", "scene")),
	}

	sections = []*codegen.SectionTemplate{
		codegen.Header(codegen.ToTitle(scriptInfo.Name)+"-Custom user text errors", "errors", imports, false),
	}

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "text-errors",
		Source: textErrorStructT,
		Data:   scriptInfo,
		FuncMap: map[string]interface{}{
			"ToTitle": codegen.ToTitle,
		},
	})

	return &codegen.File{Path: fpath, SectionTemplates: sections, IsUpdatable: true}
}

const textErrorStructT = `func main() {
	// Configuration
	cfg, err := config.New{{ ToTitle .Name }}Config()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	_ = words.LoadWordsConstants(languages.Russia, cfg.ResourcesDir)

	// Run
	app.Run(cfg)
}
`
