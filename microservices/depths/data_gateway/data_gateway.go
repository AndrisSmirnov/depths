package data_gateway

import (
	"context"
	"depths/data_gateway/mongo"
	"depths/data_gateway/redis"
)

type DataGateway struct {
	DB      *mongo.DB
	redis   *redis.Redis
	errChan chan error
}

func NewDataGateway(ctx context.Context, conf *Config) (*DataGateway, error) {
	DB, err := mongo.CreateMongoConn(
		ctx,
		conf.MongoConfig,
		conf.ErrChan,
	)
	if err != nil {
		return nil, err
	}

	redisClient, err := redis.Connect(
		ctx,
		conf.RedisConfig,
		conf.ErrChan,
	)
	if err != nil {
		return nil, err
	}

	dataGateway := &DataGateway{
		DB:      DB,
		errChan: conf.ErrChan,
		redis:   redisClient,
	}

	return dataGateway, nil
}

func (dg *DataGateway) Close() error {
	return dg.DB.CloseConnection()
}
