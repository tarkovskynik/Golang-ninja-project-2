package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	uri  string // "MONGO_URI"
	user string // "MONGO_USER"
	pass string // "MONGO_PASSWORD"
}

func NewConfig(uri, user, pass string) *Config {
	return &Config{
		uri:  uri,
		user: user,
		pass: pass,
	}
}

//return &Config{
//uri:  os.Getenv("MONGO_URI"),
//user: os.Getenv("MONGO_USER"),
//pass: os.Getenv("MONGO_PASS"),
//}

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
