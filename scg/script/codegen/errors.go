package codegen

import (
	"github.com/ThCompiler/go_game_constractor/scg/expr"
	"github.com/ThCompiler/go_game_constractor/scg/expr/scene"
	"github.com/ThCompiler/go_game_constractor/scg/generator/codegen"
	"path"
	"path/filepath"
)

// ErrorsFile returns custom textError
func ErrorsFile(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) []*codegen.File {
	errorsFile := generateTextErrors(rootPkg, rootDir, scriptInfo)

	return []*codegen.File{errorsFile}
}

func generateTextErrors(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
	var sections []*codegen.SectionTemplate

	fpath := filepath.Join(rootDir, "script", "errors", "errors.go")
	imports := []*codegen.ImportSpec{
		codegen.SCGImport(path.Join("director", "scene")),
	}

	sections = []*codegen.SectionTemplate{
		codegen.Header(codegen.ToTitle(scriptInfo.Name)+"-Custom user text errors", "errors", imports, false),
	}

	textErrors := make([]scene.Error, 0)

	for _, sc := range scriptInfo.Script {
		if sc.Error.IsText() {
			textErrors = append(textErrors, sc.Error)
		}
	}

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "text-errors",
		Source: textErrorStructT,
		Data:   textErrors,
		FuncMap: map[string]interface{}{
			"ToTitle": codegen.ToTitle,
		},
	})

	return &codegen.File{Path: fpath, SectionTemplates: sections}
}

const textErrorStructT = `{{ if . }}// text errors
var (
{{range .}}
		{{ ToTitle .Name }}Error = scene.BaseTextError{Message:"{{ .Text }}"}
{{end}}
) {{end}}
`
