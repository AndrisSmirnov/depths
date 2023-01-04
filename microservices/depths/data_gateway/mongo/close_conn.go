package mongo

import (
	"depths/pkg/log"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func (m *DB) CloseConnection() error {
	if err := m.client.Ping(m.ctx, readpref.Primary()); err == nil {
		if err := m.client.Disconnect(m.ctx); err != nil {
			return err
		}

		log.Info("Disconnected from mongoDB")
	}

	return nil
}
