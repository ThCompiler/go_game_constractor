package codegen

import (
	"path"
	"path/filepath"

	"github.com/ThCompiler/go_game_constractor/scg/expr"
	"github.com/ThCompiler/go_game_constractor/scg/generator/codegen"
)

// HandlerFile returns server file
func HandlerFile(rootPkg, rootDir string, scriptInfo expr.ScriptInfo) []*codegen.File {
	routerFile := generateRouter(rootPkg, rootDir, scriptInfo)
	routeFile := generateRoute(rootPkg, rootDir, scriptInfo)

	return []*codegen.File{routeFile, routerFile}
}

func generateRouter(_, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
	var sections []*codegen.SectionTemplate

	fpath := filepath.Join(rootDir, "internal", "controller", "http", "v1", "router.go")
	imports := []*codegen.ImportSpec{
		{Path: path.Join("net", "http")},
		{Path: path.Join("github.com", "gin-contrib", "cors")},
		{Path: path.Join("github.com", "gin-gonic", "gin")},
		codegen.SCGImport(path.Join("marusia", "runner")),
		codegen.SCGImport(path.Join("director", "scriptdirector")),
		codegen.SCGImport(path.Join("pkg", "logger")),
		codegen.SCGNamedImport(path.Join("pkg", "logger", "http"), "loghttp"),
	}

	sections = []*codegen.SectionTemplate{
		codegen.Header(codegen.ToTitle(scriptInfo.Name)+"-Router file", "v1", imports, true),
	}

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "router-file",
		Source: routerStructT,
		Data:   scriptInfo,
		FuncMap: map[string]interface{}{
			"ToTitle": codegen.ToTitle,
		},
	})

	return &codegen.File{Path: fpath, SectionTemplates: sections, IsUpdatable: true}
}

func generateRoute(_, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
	var sections []*codegen.SectionTemplate

	fpath := filepath.Join(rootDir, "internal", "controller", "http",
		"v1", codegen.SnakeCase(scriptInfo.Name)+"_handler.go")
	imports := []*codegen.ImportSpec{
		{Path: path.Join("github.com", "gin-gonic", "gin")},
		codegen.SCGImport(path.Join("marusia", "runner")),
		codegen.SCGImport(path.Join("marusia")),
		codegen.SCGImport(path.Join("marusia", "webhook")),
		codegen.SCGImport(path.Join("director", "scriptdirector")),
		codegen.SCGImport(path.Join("pkg", "logger")),
		codegen.SCGNamedImport(path.Join("pkg", "logger", "http"), "loghttp"),
	}

	sections = []*codegen.SectionTemplate{
		codegen.Header(codegen.ToTitle(scriptInfo.Name)+"-Route file", "v1", imports, true),
	}

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "route-file",
		Source: routeStructT,
		Data:   scriptInfo,
		FuncMap: map[string]interface{}{
			"ToTitle":   codegen.ToTitle,
			"SnakeCase": codegen.SnakeCase,
		},
	})

	return &codegen.File{Path: fpath, SectionTemplates: sections, IsUpdatable: true}
}

const routerStructT = `// NewRouter -.
func New{{ ToTitle .Name }}Router(
    server *gin.Engine,
    l logger.Interface,
    op{{ ToTitle .Name }} scriptdirector.SceneDirectorConfig,
    runner runner.ScriptRunner,
) {
    // Options
    corsConfig := cors.DefaultConfig()
    corsConfig.AllowOrigins = []string{"http://localhost", "https://skill-debugger.marusia.mail.ru"}
    corsConfig.AllowMethods = []string{"POST"}

    server.Use(cors.New(corsConfig))
    server.Use(gin.Recovery())
    server.Use(loghttp.GinRequestLogger(l))

    //// K8s probe
    server.GET("/healthz", func(c *gin.Context) {
        c.Status(http.StatusOK)
    })

    // Routers
    h := server.Group("/v1")
    {
        new{{ ToTitle .Name }}Route(h, op{{ ToTitle .Name }}, runner, l)
    }
}
`

const routeStructT = `type {{ ToTitle .Name }}Route struct {
    loghttp.LogObject
    sdc     scriptdirector.SceneDirectorConfig
    sRunner runner.ScriptRunner
    wh      *marusia.Webhook
}

func new{{ ToTitle .Name }}Route(handler *gin.RouterGroup, sdc scriptdirector.SceneDirectorConfig,
    sRunner runner.ScriptRunner, l logger.Interface) {
    r := &{{ ToTitle .Name }}Route{
        LogObject: loghttp.NewLogObject(l),
        sdc:       sdc,
        sRunner:   sRunner,
        wh: webhook.NewDefaultMarusiaWebhook(l, sRunner, sdc),
    }

    handler.POST("/{{ SnakeCase .Name }}", r.wh.GinHandleFunc)
}
`
