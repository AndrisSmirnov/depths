package redis

func (redisDB *Redis) GetUserConnectionTTL() {
	// 	ctx context.Context, userID string,
	// ) (*user_domain.BasicUser, error) {
	// 	userInfoByte, err := redisDB.client.Get(
	// 		ctx,
	// 		fmt.Sprintf("UserConnectionTTL-%v", userID),
	// 	).Result()
	// 	if err != nil && err == redis.Nil {
	// 		return nil, ErrNotFound
	// 	}
	//
	// 	if err != nil {
	// 		return nil, err
	// 	}
	//
	// 	basicUser := &user_domain.BasicUser{}
	// 	if err := json.Unmarshal([]byte(userInfoByte), basicUser); err != nil {
	// 		return nil, err
	// 	}
	//
	// 	return basicUser, err
}
