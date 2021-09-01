package mongo

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client interface {
	Collection(cName string) *mongo.Collection
}

type ClientImpl struct {
	Base     *mongo.Database `json:"-"`
	Hosts    string
	DBName   string
	PoolSize uint64
}

func (c *ClientImpl) SetDefaults() {
	if c.PoolSize == 0 {
		c.PoolSize = 10
	}
	if len(c.Hosts) == 0 {
		c.Hosts = "mongodb: //127.0.0.1:27017"
	}
}
func (c *ClientImpl) Init() {
	c.SetDefaults()
	c.Base = ConnectToDB(c.Hosts, c.DBName, c.PoolSize)
	err := c.Base.Client().Ping(context.Background(), nil)
	if err != nil {
		logrus.Fatalf("failed to ping: %v", err)
	}
}

func ConnectToDB(uri, dbname string, num uint64) *mongo.Database {
	ctx := context.Background()
	o := options.Client().ApplyURI(uri)
	o.SetMaxPoolSize(num)
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		logrus.Fatalf("failed to connect: %v", err)
	}

	return client.Database(dbname)
}
