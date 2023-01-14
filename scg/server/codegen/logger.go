package codegen

import (
	"path"
	"path/filepath"

	"github.com/ThCompiler/go_game_constractor/scg/expr"
	"github.com/ThCompiler/go_game_constractor/scg/generator/codegen"
)

// LoggerPrepareFile returns logger prepare file
func LoggerPrepareFile(rootPkg, rootDir string, scriptInfo expr.ScriptInfo) []*codegen.File {
	loggerPrepareIOFile := generateLoggerPrepareIO(rootPkg, rootDir, scriptInfo)
	loggerPrepareConfig := generateLoggerPrepareConfig(rootPkg, rootDir, scriptInfo)

	return []*codegen.File{loggerPrepareIOFile, loggerPrepareConfig}
}

func generateLoggerPrepareIO(_, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
	var sections []*codegen.SectionTemplate

	fpath := filepath.Join(rootDir, "pkg", "logger", "prepare", "io.go")
	imports := []*codegen.ImportSpec{
		{Path: path.Join("os")},
		{Path: path.Join("time")},
		{Path: path.Join("path", "filepath")},
		{Path: path.Join("github.com", "pkg", "errors")},
	}

	sections = []*codegen.SectionTemplate{
		codegen.Header(
			codegen.ToTitle(scriptInfo.Name)+"-Logger prepare file with open log file ",
			"prepare",
			imports,
			false,
		),
	}

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "logger-prepare-io",
		Source: loggerPrepareIOStructT,
	})

	return &codegen.File{Path: fpath, SectionTemplates: sections, IsUpdatable: false}
}

func generateLoggerPrepareConfig(_, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
	var sections []*codegen.SectionTemplate

	fpath := filepath.Join(rootDir, "pkg", "logger", "prepare", "config.go")
	imports := []*codegen.ImportSpec{
		codegen.SCGImport(path.Join("pkg", "logger")),
	}

	sections = []*codegen.SectionTemplate{
		codegen.Header(
			codegen.ToTitle(scriptInfo.Name)+"-Logger prepare file with config for logger ",
			"prepare",
			imports,
			false,
		),
	}

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "logger-prepare-config",
		Source: loggerPrepareConfigStructT,
	})

	return &codegen.File{Path: fpath, SectionTemplates: sections, IsUpdatable: false}
}

const loggerPrepareIOStructT = `const logName = "log"

func OpenLogDir(dir string) (*os.File, error) {
    if _, err := os.Stat(dir); err != nil {
        if !os.IsNotExist(err) {
            return nil, errors.Wrap(err, "error when try check log dir: ")
        }

        if err = os.MkdirAll(filepath.Dir(dir), 0755); err != nil {
            return nil, errors.Wrap(err, "error when try add log dir: ")
        }
    }

    t := time.Now().UTC()
    timeString := t.Format(time.RFC3339)
    fileName := timeString + "-" + logName + ".log"

    file, err := os.OpenFile(
        dir+"/"+fileName,
        os.O_CREATE|os.O_APPEND|os.O_WRONLY,
        0644,
    )

    if err != nil {
        return nil, errors.Wrap(err, "error when try open log file: ")
    }

    return file, nil
}
`

const loggerPrepareConfigStructT = `type Config struct {
    Level                    logger.LogLevel ` + "`" + `env-required:"true" yaml:"level" env:"LOG_LEVEL"` + "`" + `
    LogDir                   string          ` + "`" +
	`env-required:"true" yaml:"log_dir,omitempty,omitempty" env:"LOG_DIR"` + "`" + `
    UseStdAndFIle            bool            ` + "`" +
	`env-required:"true" yaml:"use_std_and_file,omitempty" env:"USER_STD_AND_FILE"` + "`" + `
    AddLowPriorityLevelToCmd bool            ` + "`" +
	`env-required:"true" yaml:"add_low_priority_level_to_cmd,omitempty" env:"ADD_LPL_TO_CMD"` + "`" + `
}

`
