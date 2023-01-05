package app

import (
	"context"
	"depths/app/domain/exchange_manager_domain"
	"depths/app/services"
	"depths/data_gateway"
)

type App struct {
	services     *services.Services
	dataGateway  *data_gateway.DataGateway
	depthManager exchange_manager_domain.IExchangeManager

	errorChan  chan error
	ctx        context.Context
	cancelFunc context.CancelFunc
}
