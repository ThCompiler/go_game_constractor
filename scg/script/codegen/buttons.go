package codegen

import (
    "github.com/ThCompiler/go_game_constractor/scg/expr"
    "github.com/ThCompiler/go_game_constractor/scg/expr/scene"
    "github.com/ThCompiler/go_game_constractor/scg/generator/codegen"
    "path"
    "path/filepath"
)

// PayloadsFile returns custom buttonsPayload
func PayloadsFile(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) []*codegen.File {
    payloadFiles := make([]*codegen.File, 0)
    for key, value := range scriptInfo.Script {
        generate := false
        for _, button := range value.Buttons {
            if button.Payload != nil {
                generate = true
                break
            }
        }
        if generate {
            payloadFiles = append(payloadFiles, generatePayloads(rootPkg, rootDir, scriptInfo.Name, key, value))
        }
    }

    return payloadFiles
}

func generatePayloads(_ string,
    rootDir string,
    scriptName string,
    sceneName string,
    sc scene.Scene,
) *codegen.File {
    var sections []*codegen.SectionTemplate

    fpath := filepath.Join(rootDir, "internal", "script", "payloads", codegen.SnakeCase(sceneName)+"_payloads.go")
    imports := []*codegen.ImportSpec{
        codegen.SCGImport(path.Join("director", "matchers")),
    }

    sections = []*codegen.SectionTemplate{
        codegen.Header(codegen.ToTitle(scriptName)+"-Custom user payloads", "payloads", imports, false),
    }

    for name, button := range sc.Buttons {
        if button.Payload != nil {
            res, _ := codegen.ConvertToGoStruct(
                button.Payload,
                codegen.ToTitle(sceneName)+codegen.ToTitle(name)+"Payload",
            )
            sections = append(sections, &codegen.SectionTemplate{
                Name:   "button-" + name + "-" + sceneName,
                Source: res,
            })
        }
    }

    return &codegen.File{Path: fpath, SectionTemplates: sections, IsUpdatable: true}
}
