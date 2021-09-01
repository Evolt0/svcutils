package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type ClientImpl struct {
	Base     *redis.Client
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
	c.Base = redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Username: c.Username,
		Password: c.Password,
		DB:       c.DB,
		PoolSize: c.PoolSize,
	})
	cmd := c.Base.Ping(context.Background())
	if err := cmd.Err(); err != nil {
		logrus.Fatalf("failed to ping: %v", err)
	}
}
