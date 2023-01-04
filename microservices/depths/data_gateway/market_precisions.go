package data_gateway

import (
	"depths/app/domain/market_precision_domain"
)

func (dg *DataGateway) InsertManyMarketPrecisions(
	data []market_precision_domain.MarketPrecision) error {
	return dg.DB.InsertMarketPrecisions(data)
}

func (dg *DataGateway) GetAllMarketPrecisions() ([]market_precision_domain.MarketPrecision, error) {
	return dg.DB.GetMarketPrecisions()
}
