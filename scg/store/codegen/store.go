package codegen

import (
	"github.com/ThCompiler/go_game_constractor/scg/expr"
	"github.com/ThCompiler/go_game_constractor/scg/generator/codegen"
	"path"
	"path/filepath"
)

// ScriptStoreFiles returns redis store for scripts implementation
func ScriptStoreFiles(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) []*codegen.File {
	repositoryFile := repository(rootPkg, rootDir, scriptInfo)
	interfaceStoreFile := scriptStore(rootPkg, rootDir, scriptInfo)

	return []*codegen.File{repositoryFile, interfaceStoreFile}
}

// client returns the files defining the gRPC client.
func repository(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
	var sections []*codegen.SectionTemplate

	fpath := filepath.Join(rootDir, "store", "redis", "repository.go")
	imports := []*codegen.ImportSpec{
		{Path: "github.com/gomodule/redigo/redis"},
		{Path: "github.com/pkg/errors"},
		{Path: path.Join(rootPkg, "pkg", "logger")},
	}

	sections = []*codegen.SectionTemplate{
		codegen.Header(codegen.ToTitle(scriptInfo.Name)+"-Redis store", "redis", imports, false),
	}

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "script-redis-repository",
		Source: scriptRedisRepositoryT,
	})

	return &codegen.File{Path: fpath, SectionTemplates: sections}
}

func scriptStore(_ string, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
	var sections []*codegen.SectionTemplate

	fpath := filepath.Join(rootDir, "store", "interface.go")
	var imports []*codegen.ImportSpec

	sections = []*codegen.SectionTemplate{
		codegen.Header(codegen.ToTitle(scriptInfo.Name)+"-Interface for script store", "repository", imports, false),
	}

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "script-store",
		Source: scriptStoreT,
	})

	return &codegen.File{Path: fpath, SectionTemplates: sections}
}

const scriptStoreT = `type ScriptStore interface {
	GetText(name string) (string, error)
	SetText(name string, value string) error
	DeleteText(name string) error
}
`

const scriptRedisRepositoryT = `
type ScriptRepository struct {
	redisPool *redis.Pool
	lg        logger.LogFunc
}

func NewScriptRepository(pool *redis.Pool, lg logger.LogFunc) *ScriptRepository {
	return &ScriptRepository{
		redisPool: pool,
		lg:        lg,
	}
}

func (repo *ScriptRepository) SetText(name string, value string) error {
	con := repo.redisPool.Get()
	defer func(con redis.Conn) {
		err := con.Close()
		if err != nil {
			repo.lg(errors.Wrapf(err, "Unsuccessful close connection to redis in set func with key: %s and value: %s",
				name, value).Error())
		}
	}(con)

	res, err := redis.String(con.Do("SET", name, value))

	if res != "OK" {
		return errors.Wrapf(err,
			"error when try add value with key: %s, and value: %s", name, value)
	}

	return nil
}

func (repo *ScriptRepository) GetText(name string) (string, error) {
	con := repo.redisPool.Get()
	defer func(con redis.Conn) {
		err := con.Close()
		if err != nil {
			repo.lg(errors.Wrapf(err, "Unsuccessful close connection to redis in get func with key: %s ",
				name).Error())
		}
	}(con)

	res, err := redis.String(con.Do("GET", name))
	if err != nil {
		return "", errors.Wrapf(err,
			"error when try get value with key: %s", name)
	}

	return res, nil
}

func (repo *ScriptRepository) DeleteText(name string) error {
	con := repo.redisPool.Get()
	defer func(con redis.Conn) {
		err := con.Close()
		if err != nil {
			repo.lg(errors.Wrapf(err, "Unsuccessful close connection to redis in delete func with key: %s ",
				name).Error())
		}
	}(con)

	_, err := redis.Int(con.Do("DEL", name))

	if err != nil {
		return errors.Wrapf(err,
			"error when try delete value with key: %s", name)
	}

	return nil
}
`
