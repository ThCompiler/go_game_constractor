package generator

// Code based on goa generator: https://github.com/goadesign/goa

import (
    "github.com/ThCompiler/go_game_constractor/scg/expr"
    "github.com/ThCompiler/go_game_constractor/scg/generator/codegen"
    "github.com/c2fo/vfs/v6"
    "golang.org/x/tools/go/packages"
    "path/filepath"
    "sort"
    "strings"

    fs "github.com/ThCompiler/go_game_constractor/scg/generator/filesystem"
)

type Param struct {
    Dir       string
    Update    bool
    AddServer bool
}

// Generate runs the code generation algorithms.
func Generate(param Param, scriptInfo expr.ScriptInfo, loc vfs.Location) (outputs []string, err1 error) {
    pkgName := strings.ToLower(codegen.CamelCase(scriptInfo.Name, false, false))
    // 1. Compute start package import path.
    var rootPkg string
    {
        base := param.Dir
        path := filepath.Join(base, codegen.Gendir)
        _, err := loc.NewLocation(path)
        if err != nil {
            return nil, err
        }

        // We create a temporary Go file to make sure the directory is a valid Go package
        dummy, err := fs.NewTempFile(loc, path, "temp.*.go")
        if err != nil {
            return nil, err
        }

        defer func() {
            if err = dummy.Delete(); err != nil {
                outputs = nil
                err1 = err
            }
        }()

        if _, err = dummy.Write([]byte("package " + pkgName)); err != nil {
            return nil, err
        }

        if err = dummy.Close(); err != nil {
            return nil, err
        }

        pkgs, err := packages.Load(&packages.Config{Mode: packages.NeedName}, path)
        if err != nil {
            return nil, err
        }
        rootPkg = pkgs[0].PkgPath
    }

    // 2. Retrieve scg generators
    genFuncs := Generators(param.AddServer)

    // 3. Generate initial set of files produced by goa code generators.
    var genFiles []*codegen.File
    for _, gen := range genFuncs {
        fs, err := gen(rootPkg, filepath.Join(codegen.Gendir), scriptInfo)
        if err != nil {
            return nil, err
        }
        genFiles = append(genFiles, fs...)
    }

    // 4. Write the files.
    written := make(map[string]struct{})
    for _, f := range genFiles {
        filename, err := f.Render(loc, param.Dir, param.Update)
        if err != nil {
            return nil, err
        }
        if filename != "" {
            written[filepath.Join(loc.Path(), filename)] = struct{}{}
        }
    }

    // 5. Compute all output filenames.
    {
        outputs = make([]string, len(written))
        cwd := loc.Path()
        i := 0
        for o := range written {
            rel, err := filepath.Rel(cwd, o)
            if err != nil {
                rel = o
            }
            outputs[i] = rel
            i++
        }
    }
    sort.Strings(outputs)

    return outputs, nil
}
