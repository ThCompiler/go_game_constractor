package codegen

import (
	"github.com/ThCompiler/go_game_constractor/scg/expr"
	"github.com/ThCompiler/go_game_constractor/scg/expr/scene"
	"github.com/ThCompiler/go_game_constractor/scg/generator/codegen"
	"path"
	"path/filepath"
)

// MatchersFile returns custom mathers
func MatchersFile(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) []*codegen.File {
	mathersFile := generateMatchers(rootPkg, rootDir, scriptInfo)

	return []*codegen.File{mathersFile}
}

func generateMatchers(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
	var sections []*codegen.SectionTemplate

	fpath := filepath.Join(rootDir, "script", "matchers", "matchers.go")
	imports := []*codegen.ImportSpec{
		codegen.SCGImport(path.Join("director", "matchers")),
	}

	sections = []*codegen.SectionTemplate{
		codegen.Header(codegen.ToTitle(scriptInfo.Name)+"-Custom user matchers", "matchers", imports, false),
	}

	regexMatchers := make([]scene.RegexMatcher, 0)
	selectsMatchers := make([]scene.SelectMatcher, 0)

	for _, sc := range scriptInfo.Script {
		for _, matcher := range sc.Matchers {
			if matcher.IsRegexMatcher() {
				mc, _ := matcher.GetRegexMatcher()
				regexMatchers = append(regexMatchers, *mc)
			}

			if matcher.IsSelectMatcher() {
				mc, _ := matcher.GetSelectsMatcher()
				selectsMatchers = append(selectsMatchers, *mc)
			}
		}
	}

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "regex-matchers",
		Source: regexMatchersStructT,
		Data:   regexMatchers,
		FuncMap: map[string]interface{}{
			"ToTitle": codegen.ToTitle,
		},
	})

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "selects-consts-matchers",
		Source: selectsConstStructT,
		Data:   selectsMatchers,
		FuncMap: map[string]interface{}{
			"ToTitle": codegen.ToTitle,
		},
	})

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "selects-matchers",
		Source: selectsMatchersStructT,
		Data:   selectsMatchers,
		FuncMap: map[string]interface{}{
			"ToTitle": codegen.ToTitle,
		},
	})

	return &codegen.File{Path: fpath, SectionTemplates: sections}
}

const regexMatchersStructT = `{{ if . }}// RegexMatchers
var (
{{range .}}
		{{ ToTitle .Name }}Matcher = matchers.NewRegexMather("{{ .Regex }}")
{{end}}
) {{end}}
`

const selectsConstStructT = `
{{ if . }}// replace string for SelectsMatchers
const (
{{range .}}
	{{ ToTitle .Name }}MatcherString = "{{ .ReplaceMessage }}"
{{end}}
) {{end}}

`

const selectsMatchersStructT = `{{ if . }}// SelectsMatchers
var (
{{range .}}
	{{ ToTitle .Name }}Matcher = matchers.NewSelectorMatcher(
	[]string{ {{ range .Selects }}
		"{{ . }}",{{end}}
	},
	{{ ToTitle .Name }}MatcherString,
	)
{{end}}
) {{end}}
`
