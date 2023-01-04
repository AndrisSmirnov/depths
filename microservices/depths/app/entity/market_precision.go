package entity

type MarketPrecision struct {
	Exchange string `bson:"_id,required"`
	Market   string `bson:"market,required"`
	// exName              string `bson:"exName,omitempty"`
	// pricePrecision      number `bson:"pricePrecision,omitempty"`
	// volumePrecision     number `bson:"volumePrecision,omitempty"`
	// minTotal            number `bson:"minTotal,omitempty"`
	// minVolume           number `bson:"minVolume,omitempty"`
	// lastPrice           number `bson:"lastPrice,omitempty"`
	// spread              number `bson:"spread,omitempty"`
	// change24            number `bson:"change24,omitempty"`
	// volume24            number `bson:"volume24,omitempty"`
	// isActive            bool   `bson:"isActive,omitempty"`
	// isFreezed           bool   `bson:"isFreezed,omitempty"`
	// step                number `bson:"step,omitempty"`
	// addressFirstSymbol  string `bson:"addressFirstSymbol,omitempty"`
	// addressSecondSymbol string `bson:"addressSecondSymbol,omitempty"`
	// addressPair         string `bson:"addressPair,omitempty"`
	// addressSecondPair   string `bson:"addressSecondPair,omitempty"`
	// startVolume         number `bson:"startVolume,omitempty"`
	// qtySteps            number `bson:"qtySteps,omitempty"`
}
