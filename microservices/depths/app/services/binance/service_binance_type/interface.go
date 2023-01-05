package service_binance_type

type IBinanceService interface {
	SubscribeMarket(marketName string) error
	UnsubscribeMarket(marketName string) error
}
