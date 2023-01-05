package app

import (
	"depths/app/domain/exchange_manager_domain"
	"depths/app/services"
	"depths/app/services/binance/service_binance"
	"depths/app/services/kuna/service_kuna"
)

func CreateServices(
	// ctx context.Context,
	// dataGateway *data_gateway.DataGateway,
	exchangeManager exchange_manager_domain.IExchangeManager,
	// errChan chan error,
) (*services.Services, error) {
	binanceService, err := service_binance.NewBinanceService(
		exchangeManager,
	)
	if err != nil {
		return nil, err
	}

	kunaService, err := service_kuna.NewKunaService(
		exchangeManager,
	)
	if err != nil {
		return nil, err
	}

	return &services.Services{
		BinanceService: binanceService,
		KunaService:    kunaService,
	}, nil
}
