package codegen

import (
    "github.com/ThCompiler/go_game_constractor/scg/expr"
    "github.com/ThCompiler/go_game_constractor/scg/expr/scene"
    "github.com/ThCompiler/go_game_constractor/scg/generator/codegen"
    "github.com/ThCompiler/go_game_constractor/scg/go/types"
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

    fpath := filepath.Join(rootDir, "internal", "script", "init.go")
    imports := []*codegen.ImportSpec{
        {Path: path.Join(rootPkg, "internal", "texts", "manager")},
        {Path: path.Join(rootPkg, "internal", "script", "scenes")},
        codegen.SCGNamedImport(path.Join("director", "scriptdirector"), "game"),
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

    fpath := filepath.Join(rootDir, "internal", "script", "scenes", "names.go")

    sections = []*codegen.SectionTemplate{
        codegen.Header(codegen.ToTitle(scriptInfo.Name)+"-Scenes Name", "scenes", nil, false),
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

type sceneLoadContext struct {
    SceneName string
    Loads     map[string]string
}

func scriptScenes(rootPkg string, rootDir string, scriptName string, sceneInfo sceneWithName) *codegen.File {
    var sections []*codegen.SectionTemplate

    fpath := filepath.Join(rootDir, "internal", "script", "scenes", codegen.SnakeCase(sceneInfo.Name)+"_scene.go")
    imports := []*codegen.ImportSpec{
        codegen.SCGImport(path.Join("director")),
        codegen.SCGImport(path.Join("director", "scriptdirector", "scene")),
        codegen.SCGImport(path.Join("scg", "go", "types")),
        codegen.SCGNamedImport(path.Join("pkg", "logger", "http"), "loghttp"),
        codegen.SCGNamedImport(path.Join("director", "scriptdirector", "matchers"), "base_matchers"),
        {Path: path.Join(rootPkg, "internal", "script", "matchers")},
        {Path: path.Join(rootPkg, "internal", "script", "errors")},
        {Path: path.Join(rootPkg, "internal", "script", "payloads")},
        {Path: path.Join(rootPkg, "internal", "texts", "manager")},
    }

    sections = []*codegen.SectionTemplate{
        codegen.Header(codegen.ToTitle(scriptName)+"-SceneStructs", "scenes", imports, true),
    }

    sections = append(sections, &codegen.SectionTemplate{
        Name:   "scene-struct-" + sceneInfo.Name,
        Source: sceneStructT,
        Data:   sceneInfo,
        FuncMap: map[string]interface{}{
            "ToTitle":                  codegen.ToTitle,
            "ConvertNameToMatcher":     matchers.ConvertNameToMatcher,
            "ConvertNameToError":       errors2.ConvertNameToError,
            "IsBaseMather":             matchers.IsCorrectNameOfMather,
            "HaveMatchedString":        haveMatchedString,
            "HaveLoadFromContextValue": haveLoadFromContextValue,
            "ToGoType":                 types.ToGoType,
        },
    })

    sceneLoadsContext := sceneLoadContext{
        SceneName: sceneInfo.Name,
        Loads:     make(map[string]string),
    }

    for _, load := range sceneInfo.Context.LoadValue {
        sceneLoadsContext.Loads[load.Name] = load.Type
    }

    for _, load := range sceneInfo.Text.Values {
        if load.FromContext != "" {
            sceneLoadsContext.Loads[load.FromContext] = load.Type
        }
    }

    sections = append(sections, &codegen.SectionTemplate{
        Name:   "load-context-scene-struct-" + sceneInfo.Name,
        Source: loadContextStructT,
        Data:   sceneLoadsContext,
        FuncMap: map[string]interface{}{
            "ToTitle":  codegen.ToTitle,
            "ToGoType": types.ToGoType,
        },
    })

    return &codegen.File{Path: fpath, SectionTemplates: sections, IsUpdatable: true}
}

func haveMatchedString(scene sceneWithName) bool {
    return len(scene.Matchers) != 0 || len(scene.Buttons) != 0
}

func haveLoadFromContextValue(values map[string]scene.Value) bool {
    found := false

    for _, val := range values {
        if val.FromContext != "" {
            found = true

            break
        }
    }

    return found
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
	loghttp.LogObject
	TextManager    manager.TextManager
	NextScene 	   SceneName
}

// React function of actions after scene has been played {{ if not .IsInfoScene }}
func (sc *{{ ToTitle .Name }}) React({{ if HaveMatchedString . }}ctx{{else}}_{{end}} *scene.Context) scene.Command { 
    {{ if .Context.SaveValue }} ctx.Set("{{ .Context.SaveValue.Name }}", types.MustConvert[{{ ToGoType .Context.SaveValue.Type }}](ctx.Request.SearchedMessage))
    
    {{end}}// TODO Write the actions after {{ ToTitle .Name }} scene has been played {{ if HaveMatchedString . }}
	switch { 
		{{ if .Buttons }}// Buttons select 
		{{end}}{{ range $name, $button := .Buttons }}case ctx.Request.NameMatched == {{ ToTitle $name }}{{ ToTitle $sceneName }}ButtonText && ctx.Request.WasButton:
            {{ if .ToScene }}sc.NextScene = {{ ToTitle .ToScene }}Scene{{end}}
		{{end}}{{ if .Matchers }}
		// Matcher select 
		{{end}}{{ range .Matchers }}case ctx.Request.NameMatched == {{ if (IsBaseMather .Name) }}base_matchers.{{ToTitle .Name}}MatchedString{{
		else}}matchers.{{ToTitle .Name}}MatchedString{{end}}:
            {{ if .ToScene }}sc.NextScene = {{ ToTitle .ToScene }}Scene{{end}}
	{{end}}default:
		sc.NextScene = {{ ToTitle .Name }}Scene
	}
	{{else}}

	sc.NextScene = {{ ToTitle .Name }}Scene // TODO: manually set next scene after reaction{{end}}
	return scene.NoCommand
}{{else}}
func (sc *{{ ToTitle .Name }}) React(_ *scene.Context) scene.Command { 
	return scene.NoCommand
}{{end}}

// Next function returning next scene
func (sc *{{ ToTitle .Name }}) Next() scene.Scene { {{ if not .IsInfoScene }}
	{{ if .NextScenes }}{{if eq (len .NextScenes) 1 }}{{ range .NextScenes }} if sc.NextScene == {{ ToTitle . }}Scene {
		return &{{ ToTitle . }}{
				TextManager: sc.TextManager,
			}
		}{{end}}{{else}}switch sc.NextScene { 
		{{ range .NextScenes }} case {{ ToTitle . }}Scene:
			return &{{ ToTitle . }}{
				TextManager: sc.TextManager,
			}
		{{end}}}{{end}}{{end}}

	return &{{ ToTitle .Name }}{
			TextManager: sc.TextManager,
	} {{ else }}{{ if .NextScene }}
	return &{{ ToTitle .NextScene }}{ {{ else }}
	return &{{ ToTitle .Name }}{ {{ end }}
			TextManager: sc.TextManager,
	} {{ end }}
}

// GetSceneInfo function returning info about scene
func (sc *{{ ToTitle .Name }}) GetSceneInfo({{ if HaveLoadFromContextValue .Text.Values }}ctx{{else}}_{{end}} *scene.Context) (scene.Info, bool) {
	{{ if .Text.Values }}var (
		{{ range $nameVar, $typeVar := .Text.Values }}{{ $nameVar }} {{ ToGoType $typeVar.Type }}{{
                if $typeVar.FromContext }} = sc.Get{{ ToTitle $typeVar.FromContext }}ContextValue(ctx){{end}}
		{{end}}
	)
	{{end}}
	// TODO Write some actions for get data for texts

	text, _ := sc.TextManager.Get{{ ToTitle .Name }}Text(
		{{ range $nameVar, $typeVar := .Text.Values }}{{$nameVar}},
		{{end}})
	return scene.Info{
		Text: text,
		ExpectedMessages: []scene.MessageMatcher{ ` + sceneMatchersStructT + ` },
		Buttons: []director.Button{ ` + sceneButtonsStructT + ` },
		{{ if .Scene.Error.IsValid }}Err: ` + sceneErrorsStructT + `,{{end}}
	}, {{ not .IsInfoScene }}
}
`

const sceneMatchersStructT = `{{ range .Matchers }}
	{{ if IsBaseMather .Name }} base_matchers.{{ ConvertNameToMatcher .Name }},{{
	else}} matchers.{{ ToTitle .Name }}Matcher,{{end}}{{end}} {{ if .Matchers }}
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

const loadContextStructT = `{{ $sceneName := .SceneName }}{{ range $nameVar, $typeVar := .Loads }}
//  Get{{ ToTitle $nameVar }}ContextValue function returning value from context about scene
func (sc *{{ ToTitle $sceneName }}) Get{{ ToTitle $nameVar }}ContextValue(ctx *scene.Context) ({{ ToGoType $typeVar }}) {
    return scene.GetContextAny[{{ ToGoType $typeVar }}](ctx, "{{$nameVar}}")
}{{ end }}
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
