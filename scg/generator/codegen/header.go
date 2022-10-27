package codegen

// Header returns a Go source file header section template.
func Header(title, pack string, imports []*ImportSpec, allowEdit bool) *SectionTemplate {
	return &SectionTemplate{
		Name:   "source-header",
		Source: headerT,
		Data: map[string]interface{}{
			"Title":       title,
			"ToolVersion": 1,
			"Pkg":         pack,
			"Imports":     imports,
			"AllowEdit":   allowEdit,
		},
	}
}

// AddImport adds imports to a section template that was generated with
// Header.
func AddImport(section *SectionTemplate, imprts ...*ImportSpec) {
	if len(imprts) == 0 {
		return
	}
	var specs []*ImportSpec
	if data, ok := section.Data.(map[string]interface{}); ok {
		if imports, ok := data["Imports"]; ok {
			specs = imports.([]*ImportSpec)
		}
		data["Imports"] = append(specs, imprts...)
	}
}

const (
	headerT = `{{if .Title}}// Code generated by scg {{.ToolVersion}}, {{if not .AllowEdit}} DO NOT EDIT {{end}}.
//
// {{.Title}}
//
// Command:
{{comment commandLine}}
//{{if not .AllowEdit}} DO NOT EDIT {{end}}.

{{end}}package {{.Pkg}}

{{if .Imports}}import {{if gt (len .Imports) 1}}(
{{end}}{{range .Imports}}	{{.Code}}
{{end}}{{if gt (len .Imports) 1}})
{{end}}
{{end}}`
)