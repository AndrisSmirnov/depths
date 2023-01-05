package service_binance

import (
	"depths/app/domain/exchange_manager_domain"
	"depths/app/services/binance/service_binance_type"
)

type binanceService struct {
	name    string
	manager exchange_manager_domain.IExchangeManagerSub
}

func (b *binanceService) GetName() string {
	return b.name
}

func NewBinanceService(
	manager exchange_manager_domain.IExchangeManagerSub,
) (
	service_binance_type.IBinanceService, error,
) {
	if manager == nil {
		return nil, ErrNilExchangeManager
	}

	b := &binanceService{
		name:    "Binance",
		manager: manager,
	}

	b.manager.RegisterManager(b)

	return b, nil
}
