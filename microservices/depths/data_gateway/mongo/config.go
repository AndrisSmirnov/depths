package mongo

type Config struct {
	user       string
	password   string
	host       string
	port       string
	database   string
	collection string
}

func NewMongoConfig(
	user string,
	password string,
	host string,
	port string,
	database string,
	collection string,
) *Config {
	return &Config{
		user:       user,
		password:   password,
		host:       host,
		port:       port,
		database:   database,
		collection: collection,
	}
}
