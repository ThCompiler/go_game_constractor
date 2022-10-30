package generator

// Code based on goa generator: https://github.com/goadesign/goa

import (
	"github.com/ThCompiler/go_game_constractor/scg/expr"
	"github.com/ThCompiler/go_game_constractor/scg/generator/codegen"
	"golang.org/x/tools/go/packages"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Generate runs the code generation algorithms.
func Generate(dir string, scriptInfo expr.ScriptInfo, update bool) (outputs []string, err1 error) {
	pkgName := strings.ToLower(codegen.CamelCase(scriptInfo.Name, false, false))
	// 1. Compute start package import path.
	var rootPkg string
	{
		base, err := filepath.Abs(dir)
		if err != nil {
			return nil, err
		}
		path := filepath.Join(base, codegen.Gendir, pkgName)
		if err = os.MkdirAll(path, 0777); err != nil {
			return nil, err
		}

		// We create a temporary Go file to make sure the directory is a valid Go package
		dummy, err := os.CreateTemp(path, "temp.*.go")
		if err != nil {
			return nil, err
		}
		defer func() {
			if err = os.Remove(dummy.Name()); err != nil {
				outputs = nil
				err1 = err
			}
		}()
		if _, err = dummy.Write([]byte("package" + pkgName)); err != nil {
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
	genfuncs := Generators(update)

	// 3. Generate initial set of files produced by goa code generators.
	var genfiles []*codegen.File
	for _, gen := range genfuncs {
		fs, err := gen(rootPkg, filepath.Join(codegen.Gendir, pkgName), scriptInfo)
		if err != nil {
			return nil, err
		}
		genfiles = append(genfiles, fs...)
	}

	// 4. Write the files.
	written := make(map[string]struct{})
	for _, f := range genfiles {
		filename, err := f.Render(dir, update)
		if err != nil {
			return nil, err
		}
		if filename != "" {
			written[filename] = struct{}{}
		}
	}

	// 5. Compute all output filenames.
	{
		outputs = make([]string, len(written))
		cwd, err := os.Getwd()
		if err != nil {
			cwd = "."
		}
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
