package codegen

import (
	"github.com/ThCompiler/go_game_constractor/scg/expr"
	"github.com/ThCompiler/go_game_constractor/scg/generator/codegen"
	"github.com/google/uuid"
	"path"
	"path/filepath"
)

// SaverStoreFile returns saver for store texts
func SaverStoreFile(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) []*codegen.File {
	constFile := constName(rootPkg, rootDir, scriptInfo)
	saverFile := saverStore(rootPkg, rootDir, scriptInfo)
	errorFile := errors(rootPkg, rootDir, scriptInfo)

	return []*codegen.File{constFile, errorFile, saverFile}
}

func saverStore(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
	var sections []*codegen.SectionTemplate

	fpath := filepath.Join(rootDir, "store", "storesaver", "init.go")
	imports := []*codegen.ImportSpec{
		{Path: path.Join(rootPkg, "store")},
		{Path: path.Join(rootPkg, "consts", "textsname"), Name: "consts"},
	}

	sections = []*codegen.SectionTemplate{
		codegen.Header(codegen.ToTitle(scriptInfo.Name)+"-Store saver", "storesaver", imports, false),
	}

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "const-key",
		Source: constInitStructT,
		Data:   codegen.ToTitle(scriptInfo.Name) + "-" + uuid.New().String(),
	})

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "check-saved-script",
		Source: checkInitStructT,
	})

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "saver-script",
		Source: saveStoreStructT,
		Data:   scriptInfo.Script,
		FuncMap: map[string]interface{}{
			"ToTitle": codegen.ToTitle,
		},
	})

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "deleter-old-scripts",
		Source: clearOldStoreStructT,
		Data:   scriptInfo,
		FuncMap: map[string]interface{}{
			"ToTitle": codegen.ToTitle,
		},
	})

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "set-script",
		Source: setStoreStructT,
	})

	return &codegen.File{Path: fpath, SectionTemplates: sections}
}

type constData struct {
	Scenes expr.Script
	Name   string
}

func constName(_ string, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
	var sections []*codegen.SectionTemplate

	fpath := filepath.Join(rootDir, "consts", "textsname", "const.go")
	var imports []*codegen.ImportSpec

	sections = []*codegen.SectionTemplate{
		codegen.Header(codegen.ToTitle(scriptInfo.Name)+"-Consts with name of scene text for", "textsname", imports, false),
	}

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "const-type-name",
		Source: constTypeNameStructT,
	})

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "const-name",
		Source: constNamesStructT,
		Data: constData{
			Scenes: scriptInfo.Script,
			Name:   codegen.ToTitle(scriptInfo.Name),
		},
		FuncMap: map[string]interface{}{
			"ToTitle": codegen.ToTitle,
		},
	})

	return &codegen.File{Path: fpath, SectionTemplates: sections}
}

func errors(_ string, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
	var sections []*codegen.SectionTemplate

	fpath := filepath.Join(rootDir, "store", "storesaver", "errors.go")
	imports := []*codegen.ImportSpec{
		{Path: "github.com/pkg/errors"},
	}

	sections = []*codegen.SectionTemplate{
		codegen.Header(codegen.ToTitle(scriptInfo.Name)+"-Error for store saver", "storesaver", imports, false),
	}

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "saver-error",
		Source: errorStructT,
	})

	return &codegen.File{Path: fpath, SectionTemplates: sections}
}

const constTypeNameStructT = `type SceneTextName string
`

const constNamesStructT = `const (
	{{ $script_name := ToTitle .Name }}{{ range $name, $scene := .Scenes }}
		// {{ ToTitle $name }}Text and {{ ToTitle $name }}TTS of text for {{ ToTitle $name }} scene
		{{ ToTitle $name }}Text = "{{ $script_name }}-{{ ToTitle $name }}Text"
		{{ ToTitle $name }}TTS = "{{ $script_name }}-{{ ToTitle $name }}TTS"
	{{ end }}
)`

const constInitStructT = `const checkKey = "{{ . }}"
`

const checkInitStructT = `
func checkScriptStore(st store.ScriptStore) bool {
	text, err := st.GetText(checkKey)
	if text != "" && err == nil {
		return true
	}
	return false
}
`

const errorStructT = `var (
	ScriptAlreadySaveError = errors.New("this script has already saved")
)
`

const saveStoreStructT = `
func saveScripts(st store.ScriptStore) error {
	var err error
{{ range $name, $scene := . }} // Set text for {{ ToTitle $name }} scene
	if err = st.SetText(consts.{{ ToTitle $name }}Text, "{{ $scene.Text.Text }}"); err != nil {
		return err
	}
	if err = st.SetText(consts.{{ ToTitle $name }}TTS, "{{ $scene.Text.TTS }}"); err != nil {
		return err
	} 

{{ end }}
	// Set info of saving text
	if err = st.SetText(checkKey, checkKey); err != nil {
		return err
	}

	return nil
}
`

const clearOldStoreStructT = `func clearOldStores(st store.ScriptStore) error {
	// Get all keys for currentScript of saving text
	keys, err := st.GetAllTextKeyForScript("{{ ToTitle .Name }}");
	if err != nil {
		return err
	}
	
	for _, key := range keys {
		_ = st.DeleteText(key)
	}

	return nil
}
`

const setStoreStructT = `
func SaveScripts(st store.ScriptStore) error {
	if checkScriptStore(st) {
		return ScriptAlreadySaveError
	}
	
	err := clearOldStores(st)
	if err != nil {
		return err
	} 
	return saveScripts(st)
}
`
