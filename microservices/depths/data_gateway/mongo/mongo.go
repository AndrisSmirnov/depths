package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	conf          *Config
	client        *mongo.Client
	marketPresCol *mongo.Collection
	errChan       chan error
	ctx           context.Context
}
