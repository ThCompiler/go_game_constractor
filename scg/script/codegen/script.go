package codegen

import (
	"github.com/ThCompiler/go_game_constractor/scg/expr"
	"github.com/ThCompiler/go_game_constractor/scg/expr/scene"
	"github.com/ThCompiler/go_game_constractor/scg/generator/codegen"
	errors2 "github.com/ThCompiler/go_game_constractor/scg/script/errors"
	"github.com/ThCompiler/go_game_constractor/scg/script/matchers"
	"path"
	"path/filepath"
)

// ScriptFile returns structs for script
func ScriptFile(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) []*codegen.File {
	directorConfigFile := directorConfig(rootPkg, rootDir, scriptInfo)
	scriptFiles := make([]*codegen.File, 0)
	for key, value := range scriptInfo.Script {
		scriptFiles = append(scriptFiles, scriptScenes(rootPkg, rootDir, scriptInfo.Name, sceneWithName{
			Scene: value,
			Name:  key,
		}))
	}
	sceneNamesFile := sceneNames(rootPkg, rootDir, scriptInfo)

	return append([]*codegen.File{directorConfigFile, sceneNamesFile}, scriptFiles...)
}

func directorConfig(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
	var sections []*codegen.SectionTemplate

	fpath := filepath.Join(rootDir, "script", "init.go")
	imports := []*codegen.ImportSpec{
		{Path: path.Join(rootPkg, "manager")},
		{Path: path.Join(rootPkg, "script", "scenes")},
		codegen.SCGNamedImport("director", "game"),
	}

	sections = []*codegen.SectionTemplate{
		codegen.Header(codegen.ToTitle(scriptInfo.Name)+"-Director config", "script", imports, false),
	}

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "director-config",
		Source: directorConfigStructT,
		Data:   scriptInfo,
		FuncMap: map[string]interface{}{
			"ToTitle": codegen.ToTitle,
			"IsLast":  storeLen,
		},
	})

	return &codegen.File{Path: fpath, SectionTemplates: sections, IsUpdatable: true}
}

func sceneNames(_ string, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
	var sections []*codegen.SectionTemplate

	fpath := filepath.Join(rootDir, "script", "scenes", "names.go")
	var imports []*codegen.ImportSpec

	sections = []*codegen.SectionTemplate{
		codegen.Header(codegen.ToTitle(scriptInfo.Name)+"-Scenes Name", "scenes", imports, false),
	}

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "scenes-name",
		Source: scenesConstantConfigStructT,
		Data:   scriptInfo,
		FuncMap: map[string]interface{}{
			"ToTitle": codegen.ToTitle,
			"ToSnake": codegen.SnakeCase,
		},
	})

	return &codegen.File{Path: fpath, SectionTemplates: sections, IsUpdatable: true}
}

type sceneWithName struct {
	scene.Scene
	Name string
}

func scriptScenes(rootPkg string, rootDir string, scriptName string, sceneInfo sceneWithName) *codegen.File {
	var sections []*codegen.SectionTemplate

	fpath := filepath.Join(rootDir, "script", "scenes", codegen.SnakeCase(sceneInfo.Name)+"_scene.go")
	imports := []*codegen.ImportSpec{
		codegen.SCGImport(path.Join("director", "scene")),
		codegen.SCGNamedImport(path.Join("director", "matchers"), "base_matchers"),
		{Path: path.Join(rootPkg, "script", "matchers")},
		{Path: path.Join(rootPkg, "script", "errors")},
		{Path: path.Join(rootPkg, "script", "payloads")},
		{Path: path.Join(rootPkg, "manager")},
	}

	sections = []*codegen.SectionTemplate{
		codegen.Header(codegen.ToTitle(scriptName)+"-SceneStructs", "scenes", imports, true),
	}

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "scene-struct-" + sceneInfo.Name,
		Source: sceneStructT,
		Data:   sceneInfo,
		FuncMap: map[string]interface{}{
			"ToTitle":              codegen.ToTitle,
			"CamelCase":            codegen.CamelCase,
			"ConvertNameToMatcher": matchers.ConvertNameToMatcher,
			"ConvertNameToError":   errors2.ConvertNameToError,
			"IsBaseMather":         matchers.IsCorrectNameOfMather,
			"HaveMatchedString":    haveMatchedString,
		},
	})

	return &codegen.File{Path: fpath, SectionTemplates: sections, IsUpdatable: true}
}

func haveMatchedString(scene sceneWithName) bool {
	return len(scene.Matchers) != 0 || len(scene.Buttons) != 0
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

const sceneStructT = `{{ $sceneName := .Name }}{{ if .Buttons }} 
const ( {{ range $name, $button := .Buttons }}
	// {{ ToTitle $name }}{{ ToTitle $sceneName }}ButtonText - text for button {{ ToTitle $name }}
	{{ ToTitle $name }}{{ ToTitle $sceneName }}ButtonText = "{{ $button.Text }}"{{end}}
)
{{end}}// {{ ToTitle .Name }} scene
type {{ ToTitle .Name }} struct {
	TextManager    manager.TextManager
	NextScene 	   SceneName
}

// React function of actions after scene has been played
func (sc *{{ ToTitle .Name }}) React({{ if HaveMatchedString . }}ctx{{else}}_{{end}} *scene.Context) scene.Command { 
	// TODO Write the actions after {{ ToTitle .Name }} scene has been played {{ if HaveMatchedString . }}
	switch { 
		{{ if .Buttons }}// Buttons select 
		{{end}}{{ range $name, $button := .Buttons }} case ctx.Request.NameMatched == {{ ToTitle $name }}{{ ToTitle $sceneName }}ButtonText && ctx.Request.WasButton:

		{{end}}{{ if .Matchers }}
		// Matcher select 
		{{end}}{{ range .Matchers }} case ctx.Request.NameMatched == {{ if (IsBaseMather .) }}base_matchers.{{ToTitle .}}MatchedString{{
		else}}matchers.{{ToTitle .}}MatchedString{{end}}:

	{{end}}}{{end}}
	

	sc.NextScene = {{ ToTitle .Name }}Scene // TODO: manually set next scene after reaction
	return scene.NoCommand
}

// Next function returning next scene
func (sc *{{ ToTitle .Name }}) Next() scene.Scene {
	{{ if .NextScenes }}switch sc.NextScene { 
		{{ range .NextScenes }} case {{ ToTitle . }}Scene:
			return &{{ ToTitle . }}{
				TextManager: sc.TextManager,
			}
		{{end}}}{{end}}

	return &{{ ToTitle .Name }}{
			TextManager: sc.TextManager,
	}
}

// GetSceneInfo function returning info about scene
func (sc *{{ ToTitle .Name }}) GetSceneInfo(_ *scene.Context) (scene.Info, bool) {
	{{ if .Text.Values }}var (
		{{range $nameVar, $typeVar := .Text.Values}}{{$nameVar}} {{$typeVar}}
		{{end}}
	)
	{{end}}
	// TODO Write some actions for get data for texts

	text, _ := sc.TextManager.Get{{ ToTitle .Name }}Text(
		{{range $nameVar, $typeVar := .Text.Values}}{{$nameVar}},
		{{end}})
	return scene.Info{
		Text: text,
		ExpectedMessages: []scene.MessageMatcher{ ` + sceneMatchersStructT + ` },
		Buttons: []scene.Button{ ` + sceneButtonsStructT + ` },
		{{ if .Scene.Error.IsValid }}Err: ` + sceneErrorsStructT + `,{{end}}
	}, true
}
`

const sceneMatchersStructT = `{{ range .Matchers }}
	{{ if IsBaseMather . }} base_matchers.{{ConvertNameToMatcher .}},{{
	else}} matchers.{{ToTitle .}}Matcher,{{end}}{{end}} {{if .Matchers}}
{{end}}`

const sceneButtonsStructT = `{{ $sceneName := .Name }}{{ range $name, $button := .Buttons }}
{ 
	Title: {{ ToTitle $name }}{{ ToTitle $sceneName}}ButtonText, {{ if $button.URL }}
	URL: "{{ $button.URL }}", {{end}} {{ if $button.Payload }}
	Payload: &payloads.{{ ToTitle $sceneName}}{{ ToTitle $name }}Payload{},{{end}}
}, {{end}} {{ if .Buttons }}
{{end}}`

const sceneErrorsStructT = `{{with .Scene.Error}}{{ if .IsBase }} base_matchers.{{ConvertNameToError .Base}}{{end}}{{
	if .IsText }} errors.{{ToTitle .Name}}Error{{end}}{{
	if .IsScene }}  scene.BaseSceneError{ Scene: &{{ToTitle .Scene}}{TextManager: sc.TextManager} }{{end}}{{end}}`

const directorConfigStructT = `
const GoodByeCommand = "{{ .GoodByeCommand }}"

func New{{ ToTitle .Name }}Script(manager  manager.TextManager) game.SceneDirectorConfig {
	return game.SceneDirectorConfig{
		StartScene:   &scenes.{{ ToTitle .StartScene }}{
						TextManager: manager,
					},
		GoodbyeScene: &scenes.{{ ToTitle .GoodByeScene }}{
						TextManager: manager,
					},
		EndCommand:   GoodByeCommand,
	}
}
`

const scenesConstantConfigStructT = `{{ if .Script }}
type SceneName string

const ( {{ range $nameScene, $scene := .Script }}
	// {{ ToTitle $nameScene }}Scene name of $nameScene
	{{ ToTitle $nameScene }}Scene = SceneName("{{ ToSnake $nameScene }}") {{end}}
	// NoScene name of next scene for ending scenes
	NoScene = SceneName("_nothing_scene")
) {{end}}
`
