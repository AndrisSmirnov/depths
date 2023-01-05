package mongo

import (
	"context"
	"depths/pkg/log"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref" // justifying it
)

func CreateMongoConn(ctx context.Context, conf *Config, errChan chan error) (*DB, error) {
	mongoDB := &DB{
		conf:    conf,
		errChan: errChan,
		ctx:     ctx,
	}

	if err := mongoDB.connectToMongo(); err != nil {
		return nil, err
	}

	mongoDB.initCollections()

	return mongoDB, nil
}

func (m *DB) connectToMongo() error {
	var (
		client *mongo.Client
		err    error
	)

	const RECONNECT_TIME_DELTA = 5

	log.Info("Connecting to MONGO...")

	for i := 1; i <= 3; i++ {
		client, err = m.connect()
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

	m.client = client

	log.Info("Connected to MONGO")

	return m.startHealthCheck()
}

func (m *DB) connect() (*mongo.Client, error) {
	connSettings := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		m.conf.user,
		m.conf.password,
		m.conf.host,
		m.conf.port,
	)

	client, err := mongo.Connect(
		m.ctx,
		options.Client().ApplyURI(connSettings))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (m *DB) initCollections() {
	m.marketPresCol = m.client.Database(m.conf.database).Collection(m.conf.collection)
}
