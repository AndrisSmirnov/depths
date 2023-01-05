package mongo

import (
	"depths/app/domain/market_precision_domain"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *DB) InsertMarketPrecisions(data []market_precision_domain.MarketPrecision) error {
	_, err := m.marketPresCol.InsertMany(m.ctx, ToSliceOfAny(data))

	return err
}

func (m *DB) GetMarketPrecisions() ([]market_precision_domain.MarketPrecision, error) {
	var (
		results []market_precision_domain.MarketPrecision
	)

	findOptions := options.Find()

	cur, err := m.marketPresCol.Find(m.ctx, bson.D{{}}, findOptions)
	if err != nil {
		return nil, err
	}

	for cur.Next(m.ctx) {
		var elem market_precision_domain.MarketPrecision
		err := cur.Decode(&elem)

		if err != nil {
			return nil, err
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return results, nil
}
