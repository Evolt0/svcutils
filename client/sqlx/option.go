package sqlx

type Option func(*ClientImpl)

func WithAddr(addr string) Option {
	return func(c *ClientImpl) {
		c.Addr = addr
	}
}

func WithPort(port int) Option {
	return func(c *ClientImpl) {
		c.Port = port
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

func WithDBName(db string) Option {
	return func(c *ClientImpl) {
		c.DBName = db
	}
}

func WithPoolSize(poolSize int) Option {
	return func(c *ClientImpl) {
		c.PoolSize = poolSize
	}
}
