package exchange_domain

type Observer interface {
	IExchange
	GetName() string
}

type IExchange interface {
	SubscribeMarket(string) error
	UnsubscribeMarket(string) error
}
