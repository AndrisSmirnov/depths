package app

import (
	"context"
	"depths/data_gateway"
	"depths/data_gateway/mongo"
	"depths/data_gateway/redis"
	"os"
	"strconv"
)

func CreateDataGateway(ctx context.Context, errChan chan error) (*data_gateway.DataGateway,
	error) {
	mongoConfig := mongo.NewMongoConfig(
		os.Getenv("MONGO_USER"),
		os.Getenv("MONGO_PASSWORD"),
		os.Getenv("MONGO_HOST"),
		os.Getenv("MONGO_PORT"),
		os.Getenv("MONGO_DATABASE"),
		os.Getenv("MONGO_COLLECTION"),
	)

	dbIndex, err := strconv.Atoi(os.Getenv("REDIS_DB_INDEX"))
	if err != nil {
		return nil, err
	}

	redisConfig := redis.NewRedisConfig(
		os.Getenv("REDIS_URL"),
		os.Getenv("REDIS_USERNAME"),
		os.Getenv("REDIS_PASSWORD"),
		dbIndex,
	)

	config := data_gateway.Config{
		MongoConfig: mongoConfig,
		RedisConfig: redisConfig,
		ErrChan:     errChan,
	}

	return data_gateway.NewDataGateway(ctx, &config)
}
