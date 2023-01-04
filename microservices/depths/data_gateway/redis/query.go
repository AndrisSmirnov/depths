package redis

import (
	"context"
	"encoding/json"
)

func (redisDB *Redis) saveUserData(ctx context.Context, key string, data any) error {
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := redisDB.client.Set(
		ctx,
		key,
		byteData,
		0,
	).Err(); err != nil {
		return err
	}

	return nil
}

func (redisDB *Redis) deleteUserData(ctx context.Context, key string) error {
	if err := redisDB.client.Del(
		ctx,
		key,
	).Err(); err != nil {
		return err
	}

	return nil
}
