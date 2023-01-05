package service_binance

import "fmt"

func (b *binanceService) SubscribeMarket(marketName string) error {
	// TODO: mock
	fmt.Printf("SUB EXCHANGE:%s:\tWITH MARKET:%s\n", b.name, marketName)

	return nil
}

func (b *binanceService) UnsubscribeMarket(marketName string) error {
	// TODO: mock
	fmt.Printf("UnSUB EXCHANGE:%s:\tWITH MARKET:%s\n", b.name, marketName)

	return nil
}
