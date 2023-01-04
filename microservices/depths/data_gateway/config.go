package data_gateway

import (
	"depths/data_gateway/mongo"
	"depths/data_gateway/redis"
)

type Config struct {
	MongoConfig *mongo.Config
	RedisConfig *redis.Config
	ErrChan     chan error
}
