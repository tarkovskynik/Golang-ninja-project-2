package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, c *Config) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(c.uri)

	opts.SetAuth(options.Credential{
		Username: c.user,
		Password: c.pass,
	})

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}
