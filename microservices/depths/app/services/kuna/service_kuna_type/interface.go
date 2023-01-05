package service_kuna_type

type IKunaService interface {
	SubscribeMarket(marketName string) error
	UnsubscribeMarket(marketName string) error
}
