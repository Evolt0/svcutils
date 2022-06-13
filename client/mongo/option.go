package mongo

type Option func(*ClientImpl)

func WithHosts(hosts string) Option {
	return func(c *ClientImpl) {
		c.Hosts = hosts
	}
}

func WithDBName(dbName string) Option {
	return func(c *ClientImpl) {
		c.DBName = dbName
	}
}

func WithPoolSize(poolSize uint64) Option {
	return func(c *ClientImpl) {
		c.PoolSize = poolSize
	}
}
