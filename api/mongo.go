package api

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Connection struct {
	Client *mongo.Client
	ctx    context.Context
}

func NewMongoConnection() (*Connection, error) {
	uri := fmt.Sprintf("mongodb://%s/%s", GetConfig("MongoHost"), GetConfig("MongoName"))

	credentials := options.Credential{
		Username: GetConfig("MongoUser"),
		Password: GetConfig("MongoPass"),
	}

	clientOpts := options.Client().ApplyURI(uri).SetAuth(credentials)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, errors.Wrap(err, "mongodb.Connect")
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, errors.Wrap(err, "testing mongo connection\n")
	}

	fmt.Println("connected to database.")
	return &Connection{
		Client: client,
		ctx:    ctx,
	}, nil
}

func (c Connection) Disconnect() {
	err := c.Client.Disconnect(c.ctx)
	if err != nil {
		return
	}
}
