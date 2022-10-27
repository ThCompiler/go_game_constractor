package codegen

import (
	"github.com/ThCompiler/go_game_constractor/scg/expr"
	"github.com/ThCompiler/go_game_constractor/scg/generator/codegen"
	"github.com/google/uuid"
	"path"
	"path/filepath"
)

// TextManagerFile returns saved text with add values from store
func TextManagerFile(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) []*codegen.File {
	usecaseFile := usecase(rootPkg, rootDir, scriptInfo)
	interfaceFile := managerInterface(rootPkg, rootDir, scriptInfo)

	return []*codegen.File{interfaceFile, usecaseFile}
}

func usecase(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
	var sections []*codegen.SectionTemplate

	fpath := filepath.Join(rootDir, "manager", "usecase", "usecase.go")
	imports := []*codegen.ImportSpec{
		{Path: path.Join(rootPkg, "store"), Name: "store"},
		{Path: path.Join(rootPkg, "pkg", "str")},
		codegen.SCGImport(path.Join("director", "scene")),
		{Path: path.Join(rootPkg, "consts", "textsname"), Name: "consts"},
	}

	sections = []*codegen.SectionTemplate{
		codegen.Header(codegen.ToTitle(scriptInfo.Name)+"-Text usecase", "usecase", imports, false),
	}

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "usecase-type",
		Source: usecaseTypeStructT,
		Data:   scriptInfo.Name + "-" + uuid.New().String(),
	})

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "usecase-func",
		Source: usecaseFuncStructT,
		Data:   scriptInfo,
		FuncMap: map[string]interface{}{
			"ToTitle": codegen.ToTitle,
			"IsLast":  storeLen,
		},
	})

	return &codegen.File{Path: fpath, SectionTemplates: sections}
}

func managerInterface(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
	var sections []*codegen.SectionTemplate

	fpath := filepath.Join(rootDir, "manager", "interface.go")
	imports := []*codegen.ImportSpec{
		codegen.SCGImport(path.Join("director", "scene")),
	}

	sections = []*codegen.SectionTemplate{
		codegen.Header(codegen.ToTitle(scriptInfo.Name)+" Interface for script text manager", "manager", imports, false),
	}

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "script-text-manger",
		Source: scriptTextManagerT,
		Data:   scriptInfo,
		FuncMap: map[string]interface{}{
			"ToTitle": codegen.ToTitle,
			"IsLast":  storeLen,
		},
	})

	return &codegen.File{Path: fpath, SectionTemplates: sections}
}

var ln = 0

func storeLen(l int) bool {
	ln++
	if l == ln {
		ln = 0
		return true
	}

	return false
}

const scriptTextManagerT = `type TextManager interface {
	{{ range $name, $scene := .Script }}
	// Get{{ ToTitle $name }}Text get text for {{$name}} scene with variables {{$lenValues := len $scene.Text.Values}}
	Get{{ ToTitle $name }}Text({{ if $lenValues }}{{
	range $nameVar, $typeVar := $scene.Text.Values}}{{$nameVar}} {{$typeVar}}{{
	if not (IsLast $lenValues) }}, {{end}}{{end}}{{end}}) (scene.Text, error)
	{{ end }}
}
`

const usecaseFuncStructT = ` {{ range $name, $scene := .Script }}
	// Get{{ ToTitle $name }}Text get text for {{$name}} scene with variables {{$lenValues := len $scene.Text.Values}}
	func (tu *TextUsecase) Get{{ ToTitle $name }}Text({{ if $lenValues }}{{
	range $nameVar, $typeVar := $scene.Text.Values}}{{$nameVar}} {{$typeVar}}{{
	if not (IsLast $lenValues) }}, {{end}}{{end}}{{end}}) (scene.Text, error) {
		text, err := tu.store.GetText(consts.{{ ToTitle $name }}Text)
		if err != nil {
			return scene.Text{}, nil
		}

		tts, err := tu.store.GetText(consts.{{ ToTitle $name }}TTS)
		if err != nil {
			return scene.Text{}, nil
		}
		
		{{ if $lenValues }} args := []interface{}{
			{{range $nameVar, $typeVar := $scene.Text.Values}}"{{$nameVar}}", {{$nameVar}},
			{{end}}
		}
		
		res := scene.Text{
			BaseText: str.StringFormat(text, args...),
			TextToSpeech: str.StringFormat(tts, args...),
		}{{ else }}
		res := scene.Text{
			BaseText: text,
			TextToSpeech: tts,
		} {{end}}

		return res, nil
	}
	{{ end }}
`

const usecaseTypeStructT = `type TextUsecase struct {
	store store.ScriptStore
}

func NewTextUsecase(store store.ScriptStore) *TextUsecase {
	return &TextUsecase{
		store: store,
	}
}
`