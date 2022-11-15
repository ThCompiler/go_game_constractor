package codegen

import (
    "github.com/ThCompiler/go_game_constractor/scg/expr"
    "github.com/ThCompiler/go_game_constractor/scg/generator/codegen"
    "path"
    "path/filepath"
)

// ConfigFile returns config file
func ConfigFile(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) []*codegen.File {
    configGoFile := generateConfigGO(rootPkg, rootDir, scriptInfo)
    configYmlFile := generateConfigYml(rootPkg, rootDir, scriptInfo)

    return []*codegen.File{configGoFile, configYmlFile}
}

func generateConfigGO(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
    var sections []*codegen.SectionTemplate

    fpath := filepath.Join(rootDir, "config", "config.go")
    imports := []*codegen.ImportSpec{
        {Path: path.Join(rootPkg, "pkg", "logger", "prepare")},
        {Path: path.Join("fmt")},
        {Path: path.Join("github.com", "ilyakaznacheev", "cleanenv")},
    }

    sections = []*codegen.SectionTemplate{
        codegen.Header(codegen.ToTitle(scriptInfo.Name)+"-Go config file ", "config", imports, true),
    }

    sections = append(sections, &codegen.SectionTemplate{
        Name:   "config-go-file",
        Source: configGoFileStructT,
        Data:   rootDir,
    })

    return &codegen.File{Path: fpath, SectionTemplates: sections, IsUpdatable: true}
}

func generateConfigYml(_ string, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
    var sections []*codegen.SectionTemplate

    fpath := filepath.Join(rootDir, "config", "config.yml")

    sections = []*codegen.SectionTemplate{}

    sections = append(sections, &codegen.SectionTemplate{
        Name:   "config-yml-file",
        Source: configYMLFileStructT,
        Data:   scriptInfo,
        FuncMap: map[string]interface{}{
            "ToTitle": codegen.ToTitle,
        },
    })

    return &codegen.File{Path: fpath, SectionTemplates: sections, IsUpdatable: true}
}

const configGoFileStructT = ` type (
    // Config -.
    Config struct {
        App   App            ` + "`" + `yaml:"app"` + "`" + `
        HTTP  HTTP           ` + "`" + `yaml:"http"` + "`" + `
        Log   prepare.Config ` + "`" + `yaml:"logger"` + "`" + `
        Redis Redis          ` + "`" + `yaml:"redis"` + "`" + `
    }

    // App -.
    App struct {
        Name         string ` + "`" + `env-required:"true" yaml:"name"    env:"APP_NAME"` + "`" + `
        Version      string ` + "`" + `env-required:"true" yaml:"version" env:"APP_VERSION"` + "`" + `
        ResourcesDir string ` + "`" + `env-required:"true" yaml:"resources_dir" env:"RESOURCES_DIR"` + "`" + `
    }

    // HTTP -.
    HTTP struct {
        Port       string ` + "`" + `env-required:"true" yaml:"port" env:"HTTP_PORT"` + "`" + `
    }

    // Redis -.
    Redis struct {
        URL string ` + "`" + `env-required:"true" yaml:"url,omitempty" env:"REDIS_URL"` + "`" + `
    }
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
    cfg := &Config{}

    err := cleanenv.ReadConfig("./{{ . }}/config/config.yml", cfg)
    if err != nil {
        return nil, fmt.Errorf("config error: %w", err)
    }

    err = cleanenv.ReadEnv(cfg)
    if err != nil {
        return nil, err
    }

    return cfg, nil
}
`

const configYMLFileStructT = `app:
  name: '{{ ToTitle .Name }}App'
  version: '1.0.0'
  resources_dir: './config/resources'

http:
  port: '8080'

logger:
  level: 'debug'
  log_dir: './app-log'
  use_std_and_file: true
  add_low_priority_level_to_cmd: true
`
