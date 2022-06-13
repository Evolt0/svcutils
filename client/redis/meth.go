package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func (c *ClientImpl) ScriptRun(
	ctx context.Context,
	script *redis.Script,
	keys []string,
	args ...interface{},
) *redis.Cmd {
	return script.Run(ctx, c.base, keys, args)
}

func (c *ClientImpl) Set(ctx context.Context, key string, data string, expireTime time.Duration) (string, error) {
	result, err := c.base.Set(ctx, key, data, expireTime).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (c *ClientImpl) Get(ctx context.Context, key string) (string, error) {
	result, err := c.base.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", err
	}
	return result, nil
}

func (c *ClientImpl) HSet(ctx context.Context, key string, field string, data string) (int64, error) {
	result, err := c.base.HSet(ctx, key, field, data).Result()
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (c *ClientImpl) HGet(ctx context.Context, key string, field string) (string, error) {
	result, err := c.base.HGet(ctx, key, field).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (c *ClientImpl) Del(ctx context.Context, key string) (int64, error) {
	result, err := c.base.Del(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (c *ClientImpl) Expire(ctx context.Context, key string, expireTime time.Duration) error {
	result, err := c.base.Expire(ctx, key, expireTime).Result()
	if err != nil {
		return err
	}
	if !result {
		return fmt.Errorf("failed to flush token expire,key is not exist")
	}
	return nil
}
