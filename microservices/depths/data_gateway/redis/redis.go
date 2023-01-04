package redis

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"depths/pkg/log"

	"github.com/go-redis/redis/v9"
)

type Redis struct {
	conf    Config
	client  *redis.Client
	ctx     context.Context
	errChan chan error
}

func Connect(
	ctx context.Context,
	conf *Config,
	errChan chan error,
) (*Redis, error) {
	redisDB := &Redis{
		conf:    *conf,
		ctx:     ctx,
		errChan: errChan,
	}

	log.Info("Connecting to REDIS...")

	if err := redisDB.connectToRedis(); err != nil {
		return nil, err
	}

	log.Info("Connected to REDIS")

	return redisDB, nil
}

func (redisDB *Redis) CheckConnection() error {
	if redisDB.client == nil {
		return ErrNotConnected
	}

	_, err := redisDB.client.Ping(redisDB.ctx).Result()

	return err
}

func (redisDB *Redis) Close() {
	log.Info("Disconnect from REDIS...")

	if redisDB.client != nil {
		redisDB.client.Close()
	}

	log.Info("Disconnected from REDIS")
}

func (redisDB *Redis) connectToRedis() error {
	var (
		client *redis.Client
		err    error
	)

	const RECONNECT_TIME_DELTA = 5

	for i := 1; i <= 3; i++ {
		client, err = redisDB.connect()
		if err == nil {
			break
		}

		log.Debug("Error connect: ", err)
		log.Debugf("Next try after %d seconds...", i*RECONNECT_TIME_DELTA)
		time.Sleep(time.Duration(i*RECONNECT_TIME_DELTA) * time.Second)
	}

	if err != nil {
		return err
	}

	if client == nil {
		return ErrConnect
	}

	redisDB.client = client

	return redisDB.startHealthCheck()
}

func (redisDB *Redis) connect() (*redis.Client, error) {
	client := redis.NewClient(
		&redis.Options{
			Addr:     redisDB.conf.ServerAddress,
			Username: redisDB.conf.Username,
			Password: redisDB.conf.Password,
			DB:       redisDB.conf.DBIndex,
		},
	)

	if _, err := client.Ping(redisDB.ctx).Result(); err != nil {
		return nil, err
	}

	return client, nil
}

func (redisDB *Redis) startHealthCheck() error {
	pingTime, err := strconv.Atoi(os.Getenv("REDIS_PING_TIME"))
	if err != nil {
		return err
	}

	ticker := time.NewTicker(time.Duration(pingTime) * time.Second)

	go func() {
		for {
			<-ticker.C

			if err := redisDB.ping(); err != nil {
				log.Warning(fmt.Sprintf("Error ping Redis: %v", err))
				log.Info("Disconnected from Redis")
				log.Info("Reconnection to Redis...")

				if err := redisDB.reconnect(); err != nil {
					log.Info("Reconnection to Redis was not successful")
				}

				if err := redisDB.ping(); err != nil {
					log.Info("Not connecter to Redis")
					log.Debug(err)
				} else {
					log.Info("Successful reconnection to Redis")
				}
			}
		}
	}()

	return nil
}

func (redisDB *Redis) ping() error {
	if redisDB.client == nil {
		return ErrNotConnected
	}

	_, err := redisDB.client.Ping(redisDB.ctx).Result()

	return err
}

func (redisDB *Redis) reconnect() error {
	client, err := redisDB.connect()
	if err != nil {
		return err
	}

	redisDB.client = client

	return nil
}
