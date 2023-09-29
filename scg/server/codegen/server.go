package codegen

import (
	"path"
	"path/filepath"

	"github.com/ThCompiler/go_game_constractor/scg/expr"
	"github.com/ThCompiler/go_game_constractor/scg/generator/codegen"
)

// ServerFile returns server pkg file
func ServerFile(rootPkg, rootDir string, scriptInfo expr.ScriptInfo) []*codegen.File {
	serverFile := generateServer(rootPkg, rootDir, scriptInfo)
	serverOptionFile := generateServerOption(rootPkg, rootDir, scriptInfo)

	return []*codegen.File{serverOptionFile, serverFile}
}

func generateServer(_, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
	var sections []*codegen.SectionTemplate

	fpath := filepath.Join(rootDir, "pkg", "httpserver", "server.go")
	imports := []*codegen.ImportSpec{
		{Path: path.Join("context")},
		{Path: path.Join("net", "http")},
		{Path: path.Join("time")},
	}

	sections = []*codegen.SectionTemplate{
		codegen.Header(codegen.ToTitle(scriptInfo.Name)+"-Http server file", "httpserver", imports, false),
	}

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "server-file",
		Source: ServerStructT,
	})

	return &codegen.File{Path: fpath, SectionTemplates: sections, IsUpdatable: false}
}

func generateServerOption(_, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
	var sections []*codegen.SectionTemplate

	fpath := filepath.Join(rootDir, "pkg", "httpserver", "option.go")
	imports := []*codegen.ImportSpec{
		{Path: path.Join("net")},
		{Path: path.Join("time")},
	}

	sections = []*codegen.SectionTemplate{
		codegen.Header(codegen.ToTitle(scriptInfo.Name)+"-Http server option file", "httpserver", imports, false),
	}

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "server-option-file",
		Source: ServerOptionStructT,
	})

	return &codegen.File{Path: fpath, SectionTemplates: sections, IsUpdatable: false}
}

const ServerStructT = `const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultAddr            = ":80"
	_defaultShutdownTimeout = 3 * time.Second
)

// Server -.
type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

// New -.
func New(server http.Handler, opts ...Option) *Server {
	httpServer := &http.Server{
		Handler:      server,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         _defaultAddr,
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
	}

	// Custom option
	for _, opt := range opts {
		opt(s)
	}

	s.start()

	return s
}

func (s *Server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
`

const ServerOptionStructT = `// Option -.
type Option func(*Server)

// Port -.
func Port(port string) Option {
	return func(s *Server) {
		s.server.Addr = net.JoinHostPort("", port)
	}
}

// ReadTimeout -.
func ReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.ReadTimeout = timeout
	}
}

// WriteTimeout -.
func WriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.WriteTimeout = timeout
	}
}

// ShutdownTimeout -.
func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}
`
