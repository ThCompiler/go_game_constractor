package generator

// Code based on goa generator: https://github.com/goadesign/goa

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/tools/go/packages"

	"github.com/ThCompiler/go_game_constractor/scg/expr"
	"github.com/ThCompiler/go_game_constractor/scg/generator/codegen"
)

// Generate runs the code generation algorithms.
func Generate(dir string, scriptInfo expr.ScriptInfo, update, addServer bool) ([]string, error) {
	pkgName := strings.ToLower(codegen.CamelCase(scriptInfo.Name, false, false))

	// 1. Compute start package import path.
	rootPkg, err := computeStartPackageImportPath(dir, pkgName)
	if err != nil {
		return nil, err
	}

	// 2. Retrieve scg generators
	genFuncs := Generators(addServer)

	// 3. Generate initial set of files produced by goa code generators.
	genFiles, errGenFiles := getGenFiles(rootPkg, scriptInfo, genFuncs)
	if errGenFiles != nil {
		return nil, errGenFiles
	}

	// 4. Write the files.
	written, errWrite := writeGeneratedFiles(genFiles, dir, update)
	if errWrite != nil {
		return nil, errWrite
	}

	// 5. Compute all output filenames.
	return computeAllOutputsFiles(written), nil
}

func computeStartPackageImportPath(dir, pkgName string) (rootPkg string, returnError error) {
	base, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}

	path := filepath.Join(base, codegen.Gendir)
	if err = os.MkdirAll(path, os.ModePerm); err != nil {
		return "", err
	}

	// We create a temporary Go file to make sure the directory is a valid Go package
	dummy, err := os.CreateTemp(path, "temp.*.go")
	if err != nil {
		return "", err
	}

	defer func() {
		if err = os.Remove(dummy.Name()); err != nil {
			returnError = err
		}
	}()

	if _, err = dummy.Write([]byte("package" + pkgName)); err != nil {
		return "", err
	}

	if err = dummy.Close(); err != nil {
		return "", err
	}

	pkgs, err := packages.Load(&packages.Config{Mode: packages.NeedName}, path)
	if err != nil {
		return "", err
	}

	return pkgs[0].PkgPath, nil
}

func getGenFiles(rootPkg string, scriptInfo expr.ScriptInfo, genFuncs []Genfunc) ([]*codegen.File, error) {
	var genFiles []*codegen.File

	for _, gen := range genFuncs {
		fs, err := gen(rootPkg, codegen.Gendir, scriptInfo)
		if err != nil {
			return nil, err
		}

		genFiles = append(genFiles, fs...)
	}

	return genFiles, nil
}

func writeGeneratedFiles(genFiles []*codegen.File, dir string, update bool) (map[string]struct{}, error) {
	written := make(map[string]struct{})

	for _, f := range genFiles {
		filename, err := f.Render(dir, update)
		if err != nil {
			return nil, err
		}

		if filename != "" {
			written[filename] = struct{}{}
		}
	}

	return written, nil
}

func computeAllOutputsFiles(written map[string]struct{}) []string {
	outputs := make([]string, len(written))

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

	sort.Strings(outputs)

	return outputs
}
