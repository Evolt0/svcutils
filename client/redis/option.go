package redis

type Option func(*ClientImpl)

func WithAddr(addr string) Option {
	return func(c *ClientImpl) {
		c.Addr = addr
	}
}

func WithPassword(password string) Option {
	return func(c *ClientImpl) {
		c.Password = password
	}
}

func WithUsername(username string) Option {
	return func(c *ClientImpl) {
		c.Username = username
	}
}

func WithDB(db int) Option {
	return func(c *ClientImpl) {
		c.DB = db
	}
}

func WithPoolSize(poolSize int) Option {
	return func(c *ClientImpl) {
		c.PoolSize = poolSize
	}
}
