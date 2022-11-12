package codegen

import (
	"github.com/ThCompiler/go_game_constractor/scg/expr"
	"github.com/ThCompiler/go_game_constractor/scg/generator/codegen"
	"path"
	"path/filepath"
)

// ConfigFile returns config file
func ConfigFile(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) []*codegen.File {
	configFile := generateConfig(rootPkg, rootDir, scriptInfo)

	return []*codegen.File{configFile}
}

func generateConfig(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
	var sections []*codegen.SectionTemplate

	fpath := filepath.Join(rootDir, "cmd", "main.go")
	imports := []*codegen.ImportSpec{
		codegen.SCGImport(path.Join("pkg", "convertor", "words")),
		codegen.SCGImport(path.Join("pkg", "convertor", "words", "languages")),
		{Path: path.Join(rootPkg, "config")},
		{Path: path.Join(rootPkg, "internal", "app")},
	}

	sections = []*codegen.SectionTemplate{
		codegen.Header(codegen.ToTitle(scriptInfo.Name)+"-Main file", "errors", imports, true),
	}

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "main-file",
		Source: configFileStructT,
		Data:   scriptInfo,
		FuncMap: map[string]interface{}{
			"ToTitle": codegen.ToTitle,
		},
	})

	return &codegen.File{Path: fpath, SectionTemplates: sections, IsUpdatable: true}
}

const configFileStructT = ` // This is your application's startup file. 
// All the basic settings take place already in the "internal/app" file. 

func main() {
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
