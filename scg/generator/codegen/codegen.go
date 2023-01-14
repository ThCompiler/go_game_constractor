package codegen

// Code based on goa generator: https://github.com/goadesign/goa

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/imports"
)

// Gendir is the name of the subdirectory of the output directory that contains
// the generated files. This directory is wiped and re-written each time goa is
// run.
const Gendir = "scg"

const (
	filePermMode = 0o644
	dirPermMode  = 0o755
)

type (
	// A File contains the logic to generate a complete file.
	File struct {
		// IsUpdatable indicates whether the file should be updated if one
		// already exists at the given path.
		IsUpdatable bool
		// SkipExist indicates whether the file should be skipped if one
		// already exists at the given path.
		SkipExist bool
		// Path returns the file path relative to the output directory.
		Path string
		// SectionTemplates is the list of file section templates in
		// order of rendering.
		SectionTemplates []*SectionTemplate
		// FinalizeFunc is called after the file has been generated. It
		// is given the absolute path to the file as argument.
		FinalizeFunc func(string) error
	}

	// A SectionTemplate is a template and accompanying render data. The
	// template format is described in the (stdlib) text/template package.
	SectionTemplate struct {
		// Name is the name reported when parsing the source fails.
		Name string
		// Source is used to create the text/template.Template that
		// renders the section text.
		Source string
		// FuncMap lists the functions used to render the templates.
		FuncMap map[string]interface{}
		// Data used as input of template.
		Data interface{}
	}
)

// Section returns the section templates with the given name or nil if not found.
func (f *File) Section(name string) []*SectionTemplate {
	var sts []*SectionTemplate

	for _, s := range f.SectionTemplates {
		if s.Name == name {
			sts = append(sts, s)
		}
	}

	return sts
}

// Render executes the file section templates and writes the resulting bytes to
// an output file. The path of the output file is computed by appending the file
// path to dir. If a file already exists with the computed path then Render
// happens the smallest integer value greater than 1 to make it unique. Renders
// returns the computed path.
func (f *File) Render(dir string, update bool) (string, error) {
	base, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}

	path := filepath.Join(base, f.Path)

	// If file already exists
	if _, err = os.Stat(path); err == nil {
		str, continueV, errPrepare := f.prepareExistsFiles(path, update)
		if !continueV {
			return str, errPrepare
		}
	}

	if err = os.MkdirAll(filepath.Dir(path), dirPermMode); err != nil {
		return "", err
	}

	err = f.baseRender(path)

	if err != nil {
		return "", err
	}

	// Format Go source files
	if filepath.Ext(path) == ".go" {
		if err = finalizeGoSource(path); err != nil {
			return "", err
		}
	}

	// Run finalizer if any
	if f.FinalizeFunc != nil {
		if err = f.FinalizeFunc(path); err != nil {
			return "", err
		}
	}

	return path, nil
}

func (f *File) prepareExistsFiles(path string, update bool) (string, bool, error) {
	if update && f.IsUpdatable {
		if err := f.renderWithUpdate(path); err != nil {
			return "", false, err
		}

		return path, false, nil
	}

	if f.SkipExist {
		return "", false, nil
	}

	if err := os.Remove(path); err != nil {
		return "", false, err
	}

	return "", true, nil
}

func (f *File) baseRender(path string) error {
	file, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		filePermMode,
	)
	if err != nil {
		return err
	}

	for _, s := range f.SectionTemplates {
		if err = s.Write(file); err != nil {
			_ = file.Close()

			return err
		}
	}

	if err = file.Close(); err != nil {
		return err
	}

	return nil
}

func (f *File) renderWithUpdate(path string) error {
	fileStrings, err := getFileLines(path)
	if err != nil {
		return err
	}

	generatedStrings, err := f.getGenLines(path)
	if err != nil {
		return err
	}

	differ := NewDiffer(generatedStrings, fileStrings)

	diff := differ.GetLineStates()

	file, err := os.OpenFile(
		path,
		os.O_WRONLY,
		filePermMode,
	)
	if err != nil {
		return err
	}

up:
	for _, df := range diff {
		if df.Tag == Equal {
			for _, line := range generatedStrings[df.firstStart:df.firstEnd] {
				if _, err = file.WriteString(line + "\n"); err != nil {
					break up
				}
			}

			continue
		}

		if df.Tag == Replace || df.Tag == Delete {
			if _, err = file.WriteString("// >>>>>>> Generated \n"); err != nil {
				break up
			}

			for _, line := range generatedStrings[df.firstStart:df.firstEnd] {
				if _, err = file.WriteString("// " + line + "\n"); err != nil {
					break up
				}
			}

			if _, err = file.WriteString("// >>>>>>> Generated \n"); err != nil {
				break up
			}
		}

		if df.Tag == Replace || df.Tag == Insert {
			for _, line := range fileStrings[df.secondStart:df.secondEnd] {
				if _, err = file.WriteString(line + "\n"); err != nil {
					break up
				}
			}
		}
	}

	if err != nil {
		_ = file.Close()

		return err
	}

	if err = file.Close(); err != nil {
		return err
	}

	return nil
}

func (f *File) getGenLines(path string) ([]string, error) {
	dummy, err := os.CreateTemp(filepath.Dir(path), "temp.*.go")
	if err != nil {
		return nil, err
	}

	for _, s := range f.SectionTemplates {
		if err = s.Write(dummy); err != nil {
			_ = dummy.Close()

			return nil, err
		}
	}

	if err = dummy.Close(); err != nil {
		return nil, err
	}

	// Format Go source files
	if err = finalizeGoSource(dummy.Name()); err != nil {
		return nil, err
	}

	res, err := getFileLines(dummy.Name())
	if err != nil {
		return nil, err
	}

	if err = os.Remove(dummy.Name()); err != nil {
		return nil, err
	}

	return res, nil
}

func getFileLines(path string) ([]string, error) {
	file, err := os.OpenFile(
		path,
		os.O_RDONLY,
		filePermMode,
	)
	if err != nil {
		return nil, err
	}

	fileStrings, err := readLines(file)
	if err != nil {
		_ = file.Close()

		return nil, err
	}

	if err = file.Close(); err != nil {
		return nil, err
	}

	return fileStrings, nil
}

func readLines(r io.Reader) ([]string, error) {
	stringLines := make([]string, 0)
	newScanner := bufio.NewScanner(r)

	for newScanner.Scan() {
		stringLines = append(stringLines, newScanner.Text())
	}

	if err := newScanner.Err(); err != nil {
		return nil, err
	}

	return stringLines, nil
}

// Write writes the section to the given writer.
func (s *SectionTemplate) Write(w io.Writer) error {
	funcs := TemplateFuncs()

	for k, v := range s.FuncMap {
		funcs[k] = v
	}

	tmpl := template.Must(template.New(s.Name).Funcs(funcs).Parse(s.Source))

	return tmpl.Execute(w, s.Data)
}

// finalizeGoSource removes unneeded imports from the given Go source file and
// runs go fmt on it.
func finalizeGoSource(path string) error {
	// Make sure file parses and print content if it does not.
	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		content, _ := os.ReadFile(path) //nolint:errcheck // there is a higher error, it does not make sense to check this

		return fmt.Errorf("%w\n========\nContent:\n%s", err, content)
	}

	// Clean unused imports
	imps := astutil.Imports(fset, file)
	for _, group := range imps {
		for _, imp := range group {
			tmpPath := strings.Trim(imp.Path.Value, `"`)

			if !astutil.UsesImport(file, tmpPath) {
				if imp.Name != nil {
					astutil.DeleteNamedImport(fset, file, imp.Name.Name, tmpPath)
				} else {
					astutil.DeleteImport(fset, file, tmpPath)
				}
			}
		}
	}

	ast.SortImports(fset, file)

	w, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}

	if err = format.Node(w, fset, file); err != nil {
		return err
	}

	w.Close()

	// Format code using goimport standard
	bs, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	opt := imports.Options{
		Comments:   true,
		FormatOnly: true,
	}

	bs, err = imports.Process(path, bs, &opt)
	if err != nil {
		return err
	}

	return os.WriteFile(path, bs, os.ModePerm)
}
