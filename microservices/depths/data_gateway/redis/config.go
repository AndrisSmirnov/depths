package redis

type Config struct {
	ServerAddress string
	Username      string
	Password      string
	DBIndex       int
}

func NewRedisConfig(
	serverAddress,
	userName,
	password string,
	dbIndex int,
) *Config {
	return &Config{
		ServerAddress: serverAddress,
		Username:      userName,
		Password:      password,
		DBIndex:       dbIndex,
	}
}
