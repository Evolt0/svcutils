package mongo

import "go.mongodb.org/mongo-driver/mongo"

func (c *ClientImpl) Collection(cName string) *mongo.Collection {
	return c.Base.Collection(cName)
}
