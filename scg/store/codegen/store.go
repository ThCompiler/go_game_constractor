package codegen

import (
    "github.com/ThCompiler/go_game_constractor/scg/expr"
    "github.com/ThCompiler/go_game_constractor/scg/generator/codegen"
    "path/filepath"
)

// ScriptStoreFiles returns redis store for scripts implementation
func ScriptStoreFiles(rootPkg string, rootDir string, scriptInfo expr.ScriptInfo) []*codegen.File {
    repositoryFile := repository(rootPkg, rootDir, scriptInfo)
    interfaceStoreFile := scriptStore(rootPkg, rootDir, scriptInfo)

    return []*codegen.File{repositoryFile, interfaceStoreFile}
}

// client returns the files defining the gRPC client.
func repository(_ string, rootDir string, scriptInfo expr.ScriptInfo) *codegen.File {
    var sections []*codegen.SectionTemplate

    fpath := filepath.Join(rootDir, "internal", "texts", "store", "redis", "repository.go")
    imports := []*codegen.ImportSpec{
        {Path: "github.com/go-redis/redis/v8", Name: "redis"},
        {Path: "context"},
        {Path: "github.com/pkg/errors"},
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

    fpath := filepath.Join(rootDir, "internal", "texts", "store", "interface.go")
    var imports []*codegen.ImportSpec

    sections = []*codegen.SectionTemplate{
        codegen.Header(codegen.ToTitle(scriptInfo.Name)+"-Interface for script store", "store", imports, false),
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
	GetAllTextKeyForScript(name string) ([]string, error)
}
`

const scriptRedisRepositoryT = `type ScriptRepository struct {
	rdc *redis.Client
	ctx context.Context
}

func NewScriptRepository(RedisClient *redis.Client) *ScriptRepository {
	return &ScriptRepository{
		rdc: RedisClient,
		ctx: context.Background(),
	}
}

func (repo *ScriptRepository) SetText(name string, value string) error {
	err := repo.rdc.Set(repo.ctx, name, value, 0).Err()

	if err != nil {
		return errors.Wrapf(err,
			"error when try add value with key: %s, and value: %s", name, value)
	}

	return nil
}

func (repo *ScriptRepository) GetText(name string) (string, error) {
	res, err := repo.rdc.Get(repo.ctx, name).Result()
	if err != nil {
		return "", errors.Wrapf(err,
			"error when try get value with key: %s", name)
	}

	return res, nil
}

func (repo *ScriptRepository) DeleteText(name string) error {
	err := repo.rdc.Del(repo.ctx, name).Err()

	if err != nil {
		return errors.Wrapf(err,
			"error when try delete value with key: %s", name)
	}

	return nil
}

func (repo *ScriptRepository) GetAllTextKeyForScript(name string) ([]string, error) {
	res := make([]string, 0)
	iter := repo.rdc.Scan(repo.ctx, 0, name+"-*", 0).Iterator()
	for iter.Next(repo.ctx) {
		res = append(res, iter.Val())
	}

	if err := iter.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

`
