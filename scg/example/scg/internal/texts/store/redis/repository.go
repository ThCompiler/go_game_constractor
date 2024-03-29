// Code generated by scg 1,  DO NOT EDIT .
//
// EchoGame-Redis store
//
// Command:
// scg
// DO NOT EDIT .

package redis

import (
	"context"

	redis "github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type ScriptRepository struct {
	rdc *redis.Client
	ctx context.Context
}

func NewScriptRepository(RedisClient *redis.Client) *ScriptRepository {
	return &ScriptRepository{
		rdc: RedisClient,
		ctx: context.Background(),
	}
}

func (repo *ScriptRepository) SetText(name, value string) error {
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
