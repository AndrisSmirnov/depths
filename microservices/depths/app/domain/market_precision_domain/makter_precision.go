package market_precision_domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type MarketPrecision struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Exchange  string             `json:"exchange",bson:"exchange,required"`
	Market    string             `json:"market",bson:"market,required"`
	ExName    string             `json:"exName",bson:"exName,omitempty"`
	IsActive  bool               `json:"isActive",bson:"isActive,omitempty"`
	IsFreezed bool               `json:"isFreezed",bson:"isFreezed,omitempty"`
}
