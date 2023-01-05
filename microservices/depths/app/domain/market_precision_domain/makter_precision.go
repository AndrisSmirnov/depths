package market_precision_domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type MarketPrecision struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Exchange  string             `bson:"exchange,required"`
	Market    string             `bson:"market,required"`
	ExName    string             `bson:"exName,omitempty"`
	IsActive  bool               `bson:"isActive,omitempty"`
	IsFreezed bool               `bson:"isFreezed,omitempty"`
}

func (m *MarketPrecision) IsActiveMarket() bool {
	return m.IsActive
}

func (m *MarketPrecision) IsFrozenMarket() bool {
	return m.IsFreezed
}
