package mongo

import (
	"depths/pkg/log"
	"fmt"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func (m *DB) startHealthCheck() error {
	pingTime, err := strconv.Atoi(os.Getenv("MONGO_PING_TIME"))
	if err != nil {
		return err
	}

	ticker := time.NewTicker(time.Duration(pingTime) * time.Second)

	go func() {
		for {
			<-ticker.C

			if err := m.ping(); err != nil {
				log.Warning(fmt.Sprintf("Error ping MongoDB: %v", err))
				log.Info("Disconnected from MongoDB")
				log.Info("Reconnection to MongoDB...")

				if err := m.reconnect(); err != nil {
					log.Info("Reconnection to MongoDB was not successful")
				}

				if err := m.ping(); err != nil {
					log.Info("Not connecter to MongoDB")
					log.Debug(err)
				} else {
					log.Info("Successful reconnection to MongoDB")
				}
			}
		}
	}()

	return nil
}

func (m *DB) ping() error {
	if m.client == nil {
		return ErrNotConnected
	}

	return m.client.Ping(m.ctx, &readpref.ReadPref{})
}

func (m *DB) reconnect() error {
	client, err := m.connect()
	if err != nil {
		return err
	}

	m.client = client

	return nil
}
