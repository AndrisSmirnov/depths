package exchange_manager_domain

import (
	"depths/app/domain/market_precision_domain"
	"fmt"
)

func (a ActiveExchangeMarkets) isInExchange(
	market market_precision_domain.MarketPrecision) bool {
	if _, ok := a[market.Exchange]; !ok {
		fmt.Printf("\tnot found a[%s]\n", market.Exchange)
		return false
	}

	if _, ok := a[market.Exchange][market.Market]; !ok {
		fmt.Printf("\tnot found a[%s][%s]\n", market.Exchange, market.Market)
		return false
	}

	delete(a[market.Exchange], market.Market)

	return true
}
