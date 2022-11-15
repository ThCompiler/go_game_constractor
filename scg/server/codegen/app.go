package codegen

import (
    "github.com/ThCompiler/go_game_constractor/scg/expr"
    "github.com/ThCompiler/go_game_constractor/scg/generator/codegen"
    "path"
    "path/filepath"
    "strings"
)

// AppFile returns app file
func AppFile(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) []*codegen.File {
    appFile := generateApp(rootPkg, rootDir, scriptInfo)

    return []*codegen.File{appFile}
}

func generateApp(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
    var sections []*codegen.SectionTemplate

    fpath := filepath.Join(rootDir, "internal", strings.ToLower(codegen.ToTitle(scriptInfo.Name)), "app.go")
    imports := []*codegen.ImportSpec{
        {Path: path.Join("fmt")},
        {Path: path.Join("os")},
        {Path: path.Join("os", "signal")},
        {Path: path.Join("io")},
        {Path: path.Join("log")},
        {Path: path.Join("syscall")},
        {Path: path.Join(rootPkg, "internal", "script")},
        {Path: path.Join(rootPkg, "internal", "controller", "http", "v1")},
        {Path: path.Join(rootPkg, "internal", "texts", "manager", "usecase")},
        {Path: path.Join(rootPkg, "internal", "texts", "store", "redis"), Name: codegen.SnakeCase(scriptInfo.Name) + "_redis"},
        {Path: path.Join(rootPkg, "internal", "texts", "store", "storesaver")},
        {Path: path.Join(rootPkg, "pkg", "logger", "prepare")},
        {Path: path.Join(rootPkg, "pkg", "httpserver")},
        {Path: path.Join(rootPkg, "config")},
        codegen.SCGImport(path.Join("marusia", "runner", "hub")),
        codegen.SCGImport(path.Join("pkg", "logger", "zap")),
        {Path: path.Join("github.com", "go-redis", "redis", "v8"), Name: "redis"},
        {Path: path.Join("github.com", "gin-gonic", "gin")},
    }
    sections = []*codegen.SectionTemplate{
        codegen.Header(codegen.ToTitle(scriptInfo.Name)+"-App file", strings.ToLower(codegen.ToTitle(scriptInfo.Name)), imports, true),
    }

    sections = append(sections, &codegen.SectionTemplate{
        Name:   "app-file",
        Source: AppStructT,
        Data:   scriptInfo,
        FuncMap: map[string]interface{}{
            "ToTitle":   codegen.ToTitle,
            "CamelCase": codegen.LowerCamelCase,
            "SnakeCase": codegen.SnakeCase,
        },
    })

    return &codegen.File{Path: fpath, SectionTemplates: sections, IsUpdatable: true}
}

const AppStructT = `// Run creates objects via constructors.
func Run(cfg *config.Config) {
    // Logger
    var logOut io.Writer

    if cfg.Log.LogDir != "" {
        file, err := prepare.OpenLogDir(cfg.Log.LogDir)
        if err != nil {
            log.Fatalf("Create logger error: %s", err)
        }

        defer func() {
            err = file.Close()
            log.Fatalf("Close log file error: %s", err)
        }()

        logOut = file
    } else {
        logOut = os.Stderr
    }

    l := zap.New(
        zap.Params{
            AppName:                  cfg.App.Name,
            LogDir:                   cfg.Log.LogDir,
            Level:                    cfg.Log.Level,
            UseStdAndFIle:            cfg.Log.UseStdAndFIle,
            AddLowPriorityLevelToCmd: cfg.Log.AddLowPriorityLevelToCmd,
        },
        logOut,
    )

    defer func() {
        _ = l.Sync()
    }()
    

    // Redis
    opt, err := redis.ParseURL(cfg.Redis.URL)
    if err != nil {
        l.Fatal(fmt.Errorf("app - Run - redis - redis.New: %w", err))
    }
    rdb := redis.NewClient(opt)

    
    // Repository
    {{ CamelCase .Name }}Store := {{ SnakeCase .Name }}_redis.NewScriptRepository(rdb)

    err = storesaver.SaveScripts({{ CamelCase .Name }}Store)
    if err != nil && err != storesaver.ScriptAlreadySaveError {
        l.Fatal(fmt.Errorf("app - Run - store - saver.SaveStore: %w", err))
    }
    

    // Game Director
    gameDirectorConfig{{ ToTitle .Name }} := script.New{{ ToTitle .Name }}Script(usecase.NewTextUsecase({{ CamelCase .Name }}Store))


    // HTTP Server
    runner := hub.NewHub()

    appHandler := gin.New()
    v1.New{{ ToTitle .Name }}Router(appHandler, l, gameDirectorConfig{{ ToTitle .Name }}, runner)
    httpServer := httpserver.New(appHandler, httpserver.Port(cfg.HTTP.Port))

    // Waiting signal
    interrupt := make(chan os.Signal, 1)
    signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

    go runner.Run()

    select {
    case s := <-interrupt:
        l.Info("app - Run - signal: " + s.String())
    case err := <-httpServer.Notify():
        l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
    }

    runner.StopHub()

    // Shutdown
    err = httpServer.Shutdown()
    if err != nil {
        l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
    }
}
`
