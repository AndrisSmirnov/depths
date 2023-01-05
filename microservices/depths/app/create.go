package app

import (
	"context"
	"depths/app/domain/exchange_manager_domain"
)

func CreateApplication(ctx context.Context, cancelFunc func()) (*App, error) {
	const ERROR_AMOUNT = 5

	errChan := make(chan error, ERROR_AMOUNT)

	dataGateway, err := CreateDataGateway(ctx, errChan)
	if err != nil {
		return nil, err
	}

	exchangeManager, err := exchange_manager_domain.NewExchangeManager(dataGateway)
	if err != nil {
		return nil, err
	}

	services, err := CreateServices(exchangeManager)
	if err != nil {
		return nil, err
	}

	return &App{
		dataGateway:  dataGateway,
		depthManager: exchangeManager,
		services:     services,
		cancelFunc:   cancelFunc,
	}, nil
}
