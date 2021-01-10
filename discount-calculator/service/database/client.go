package database

import (
	"context"
	"fmt"
	"github.com/iurybranco/backend-challenge/discount-calculator/service/database/documents"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ErrNoDocuments = mongo.ErrNoDocuments

type client struct {
	dbClient          *mongo.Client
	userCollection    *mongo.Collection
	productCollection *mongo.Collection
}

func New(config Config) (Client, error) {
	dbClient, err := mongo.Connect(nil, options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s:%d", config.Host, config.Port)).
		SetAuth(options.Credential{
			Username: config.Username,
			Password: config.Password,
		}),
	)
	if err != nil {
		return nil, err
	}
	db := dbClient.Database(config.Database)
	return &client{
		dbClient:          dbClient,
		userCollection:    db.Collection(config.UserCollection),
		productCollection: db.Collection(config.ProductCollection),
	}, nil
}

func (c *client) GetUser(id int32) (*documents.User, error) {
	result := c.userCollection.FindOne(context.Background(), bson.M{"_id": id})
	var user documents.User
	if err := result.Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *client) GetProduct(id int32) (*documents.Product, error) {
	result := c.productCollection.FindOne(context.Background(), bson.M{"_id": id})
	var product documents.Product
	if err := result.Decode(&product); err != nil {
		return nil, err
	}
	return &product, nil
}

func (d *client) Close() error {
	return d.dbClient.Disconnect(context.Background())
}
