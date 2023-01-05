package service_kuna

import "fmt"

func (k *kunaService) SubscribeMarket(marketName string) error {
	// TODO: mock
	fmt.Printf("SUB EXCHANGE:%s:\t\tWITH MARKET:%s\n", k.name, marketName)

	return nil
}

func (k *kunaService) UnsubscribeMarket(marketName string) error {
	// TODO: mock
	fmt.Printf("UnSUB EXCHANGE:%s:\t\tWITH MARKET:%s\n", k.name, marketName)

	return nil
}
