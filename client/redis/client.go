package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"time"
)

type Client interface {
	ScriptRun(ctx context.Context, script *redis.Script, keys []string, args ...interface{}) *redis.Cmd
	Set(ctx context.Context, key string, data string, expireTime time.Duration) (string, error)
	Get(ctx context.Context, key string) (string, error)
	HSet(ctx context.Context, key string, field string, data string) (int64, error)
	HGet(ctx context.Context, key string, field string) (string, error)
	Del(ctx context.Context, key string) (int64, error)
	Expire(ctx context.Context, key string, expireTime time.Duration) error
}

type ClientImpl struct {
	base     *redis.Client
	Addr     string
	Password string
	Username string
	DB       int
	PoolSize int
}

func NewClientImpl(opts ...Option) (*ClientImpl, error) {
	c := &ClientImpl{}
	c.SetDefaults()
	for _, opt := range opts {
		opt(c)
	}
	c.Init()
	return c, nil
}

func (c *ClientImpl) SetDefaults() {
}

func (c *ClientImpl) Init() {
	c.base = redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Username: c.Username,
		Password: c.Password,
		DB:       c.DB,
		PoolSize: c.PoolSize,
	})
	cmd := c.base.Ping(context.Background())
	if err := cmd.Err(); err != nil {
		logrus.Fatalf("failed to ping: %v", err)
	}
}
