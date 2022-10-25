package codegen

import (
	"github.com/ThCompiler/go_game_constractor/scg/expr"
	"github.com/ThCompiler/go_game_constractor/scg/expr/scene"
	"github.com/ThCompiler/go_game_constractor/scg/generator/codegen"
	"path"
	"path/filepath"
)

// ScriptFile returns saved text with add values from store
func ScriptFile(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) []*codegen.File {
	directorConfigFile := directorConfig(rootPkg, rootDir, scriptInfo)
	scriptFile := scriptScenes(rootPkg, rootDir, scriptInfo)

	return []*codegen.File{scriptFile, directorConfigFile}
}

func directorConfig(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
	var sections []*codegen.SectionTemplate

	fpath := filepath.Join(rootDir, "script", "init.go")
	imports := []*codegen.ImportSpec{
		{Path: path.Join(rootPkg, "manager")},
		codegen.SCGNamedImport("director", "game"),
	}

	sections = []*codegen.SectionTemplate{
		codegen.Header(codegen.ToTitle(scriptInfo.Name)+"-Director config", "script", imports, false),
	}

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "director-config",
		Source: directorConfigStruct,
		Data:   scriptInfo,
		FuncMap: map[string]interface{}{
			"ToTitle": codegen.ToTitle,
			"IsLast":  storeLen,
		},
	})

	return &codegen.File{Path: fpath, SectionTemplates: sections}
}

type sceneWithName struct {
	scene.Scene
	Name string
}

func scriptScenes(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
	var sections []*codegen.SectionTemplate

	fpath := filepath.Join(rootDir, "script", "script.go")
	imports := []*codegen.ImportSpec{
		codegen.SCGImport(path.Join("director", "scene")),
		{Path: path.Join(rootPkg, "manager")},
	}

	sections = []*codegen.SectionTemplate{
		codegen.Header(codegen.ToTitle(scriptInfo.Name)+"-SceneStructs", "script", imports, true),
	}

	for key, value := range scriptInfo.Script {
		sc := sceneWithName{
			Scene: value,
			Name:  key,
		}
		sections = append(sections, &codegen.SectionTemplate{
			Name:   "scene-struct-" + key,
			Source: sceneStructT,
			Data:   sc,
			FuncMap: map[string]interface{}{
				"ToTitle":   codegen.ToTitle,
				"CamelCase": codegen.CamelCase,
			},
		})
	}

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "scene-struct-goodbye",
		Source: sceneStructT,
		Data:   scriptInfo.GoodByeScene,
		FuncMap: map[string]interface{}{
			"ToTitle":   codegen.ToTitle,
			"CamelCase": codegen.CamelCase,
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

const sceneStructT = `type {{ ToTitle .Name }} struct {
		TextManager    manager.TextManager
	}
	
	func (sc *{{ ToTitle .Name }}) React(_ *scene.Context) scene.Command {
		// TODO
		return scene.NoCommand
	}
	
	func (sc *{{ ToTitle .Name }}) Next() scene.Scene {
		//TODO
		return &{{ ToTitle .Name }}{TextManager: sc.TextManager}
	}
	
	func (sc *{{ ToTitle .Name }}) GetSceneInfo(ctx *scene.Context) (scene.Info, bool) {
		//TODO
		
		var (
			{{range $nameVar, $typeVar := .Text.Values}}{{$nameVar}} {{$typeVar}}
			{{end}}
		)
		
		text, _ := sc.TextManager.Get{{ ToTitle .Name }}Text(
			{{range $nameVar, $typeVar := .Text.Values}}{{$nameVar}},
			{{end}})
		return scene.Info{
			Text: text,
			ExpectedMessages: []scene.MessageMatcher{},
			Buttons: []scene.Button{},
			Err: &scene.BaseSceneError{},
		}, true
	}
`

const directorConfigStruct = `
const GoodByeCommand = "{{ .GoodByeCommand }}"

func New{{ ToTitle .Name }}Script(manager  manager.TextManager) game.SceneDirectorConfig {
	return game.SceneDirectorConfig{
		StartScene:   &{{ ToTitle .StartScene }}{manager},
		GoodbyeScene: &{{ ToTitle .GoodByeScene.Name }}{manager},
		EndCommand:   GoodByeCommand,
	}
}
`
