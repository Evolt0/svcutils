package sqlx

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Client interface {
}

type ClientImpl struct {
	Base     *sqlx.DB
	Addr     string
	Port     int
	Password string
	Username string
	DBName   string
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
	if c.Addr == "" {
		c.Addr = "127.0.0.1"
	}

	if c.Port == 0 {
		c.Port = 3306
	}

	if c.PoolSize == 0 {
		c.PoolSize = 10
	}
}

func (c *ClientImpl) Init() {
	c.SetDefaults()
	var err error
	c.Base, err = sqlx.Open("mysql", c.parse())
	if err != nil {
		logrus.Fatalf("failed to open: %v", err)
	}
	c.Base.SetMaxOpenConns(c.PoolSize)
	c.Base.SetMaxIdleConns(c.PoolSize / 2)
	if err := c.Base.Ping(); err != nil {
		logrus.Fatalf("failed to ping: %v", err)
	}
}

func (c *ClientImpl) parse() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", c.Username, c.Password, c.Addr, c.Port, c.DBName)
}
