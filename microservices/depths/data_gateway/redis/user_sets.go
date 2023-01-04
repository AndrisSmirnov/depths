package redis

import "context"

func (redisDB *Redis) GetUserFromRedisSets(ctx context.Context, userID string) error {
	exist, err := redisDB.client.SIsMember(
		ctx,
		"UserSets",
		userID,
	).Result()

	if !exist {
		return ErrNotFound
	}

	return err
}
