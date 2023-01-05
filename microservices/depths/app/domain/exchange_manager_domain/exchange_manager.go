package exchange_manager_domain

import (
	"depths/app/domain/exchange_domain"
	"depths/app/domain/market_precision_domain"
	"depths/pkg/log"
	"depths/voc"
	"fmt"
	"os"
	"strconv"
	"time"
)

type IExchangeManagerSub interface {
	RegisterManager(o exchange_domain.Observer)
	Deregister(o exchange_domain.Observer)
}

type IExchangeManager interface {
	IExchangeManagerSub
	Start() error
}

type (
	ActiveMarket          map[string]interface{}
	ActiveExchangeMarkets map[string]ActiveMarket
	ObserverMap           map[string]exchange_domain.Observer

	ExchangeManager struct {
		observerMap ObserverMap
		Active      ActiveExchangeMarkets
		dg          IDataGateway
		i           int
	}
)

func NewExchangeManager(dg IDataGateway) (IExchangeManager, error) {
	if dg == nil {
		return nil, errNilDataGateway
	}

	return &ExchangeManager{
		dg:          dg,
		observerMap: make(ObserverMap),
		Active:      make(ActiveExchangeMarkets),
	}, nil
}

func (d *ExchangeManager) RegisterManager(o exchange_domain.Observer) {
	d.observerMap[o.GetName()] = o
}

func (d *ExchangeManager) Deregister(o exchange_domain.Observer) {
	delete(d.observerMap, o.GetName())
}

func (d *ExchangeManager) Start() error {
	updateExchangeMarketsTime, err := strconv.Atoi(os.Getenv("MANAGER_TIME"))
	if err != nil {
		return err
	}

	ticker := time.NewTicker(time.Duration(updateExchangeMarketsTime) * time.Second)

	for range ticker.C {
		markets, err := d.dg.GetAllMarketPrecisions()
		if err != nil {
			log.Warning(
				fmt.Sprintf("%s with err:%v", voc.ErrorWithGetMongoMarkets, err))
		}

		d.searchInMarkets(markets)
	}

	return nil
}

func (d *ExchangeManager) searchInMarkets(markets []market_precision_domain.MarketPrecision) {
	fmt.Printf("\t\td.Active:\t%v\n", d.Active)

	newActiveExchangeMarkets := make(ActiveExchangeMarkets)

	for _, m := range markets {
		if m.IsActive && !m.IsFreezed {
			if _, ok := newActiveExchangeMarkets[m.Exchange]; !ok {
				newActiveExchangeMarkets[m.Exchange] = make(ActiveMarket)
			}

			newActiveExchangeMarkets[m.Exchange][m.Market] = nil

			if !d.Active.isInExchange(m) {
				fmt.Printf("ADD\nEXCHANGE\t:%v\nMARKET\t:%v\n", m.Exchange, m.Market)
				// TODO: add open ws EXCHANGE & MARKET
				if err := d.observerMap.find(m.Exchange).SubscribeMarket(m.Market); err != nil {
					log.Warning(
						fmt.Sprintf("%s:%s -> %s",
							voc.ErrorSubscribe,
							m.Exchange,
							m.Market,
						))
				}
			}
		}
	}

	for exchange, activeMarkets := range d.Active {
		for market := range activeMarkets {
			fmt.Printf("RM\nEXCHANGE\t:%v\nMARKET\t:%v\n", exchange, market)
			// TODO: add close ws EXCHANGE & MARKET
			if err := d.observerMap.find(exchange).UnsubscribeMarket(market); err != nil {
				log.Warning(
					fmt.Sprintf("%s:%s -> %s",
						voc.ErrorUnsubscribe,
						exchange,
						market,
					))
			}
		}
	}

	d.Active = newActiveExchangeMarkets

	d.i++
	fmt.Printf("\t\tnewActive:\t%v\n"+
		"-------------------------------------------------"+
		"-------------------------------------------------  %d\n", newActiveExchangeMarkets,
		d.i)
}
