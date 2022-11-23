package codegen

import (
	"github.com/ThCompiler/go_game_constractor/scg/expr"
	"github.com/ThCompiler/go_game_constractor/scg/generator/codegen"
	"path/filepath"
)

// PkgFiles returns files for pkg dir
func PkgFiles(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) []*codegen.File {
	stringFormatFile := stringFormat(rootPkg, rootDir, scriptInfo)

	return []*codegen.File{stringFormatFile}
}

func stringFormat(_ string, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
	var sections []*codegen.SectionTemplate

	fpath := filepath.Join(rootDir, "pkg", "str", "string_format.go")
	imports := []*codegen.ImportSpec{
		{Path: "strings"},
		{Path: "fmt"},
	}

	sections = []*codegen.SectionTemplate{
		codegen.Header(codegen.ToTitle(scriptInfo.Name)+"-Set values to string in tags {valueName}", "str", imports, false),
	}

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "string-format",
		Source: stringFormatStructT,
	})

	return &codegen.File{Path: fpath, SectionTemplates: sections}
}

const stringFormatStructT = `// StringFormat set values to string in tags {valueName}
func StringFormat(format string, args ...interface{}) string {
	preparedArgs := make([]string, len(args))
	for i, v := range args {
		if i%2 == 0 {
			preparedArgs[i] = fmt.Sprintf("{%v}", v)
		} else {
			preparedArgs[i] = fmt.Sprint(v)
		}
	}
	r := strings.NewReplacer(preparedArgs...)
	return fmt.Sprint(r.Replace(format))
}
`
