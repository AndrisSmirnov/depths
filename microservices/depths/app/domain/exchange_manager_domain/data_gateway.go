package exchange_manager_domain

import "depths/app/domain/market_precision_domain"

type IDataGateway interface {
	GetAllMarketPrecisions() ([]market_precision_domain.MarketPrecision, error)
}
