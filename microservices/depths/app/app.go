package app

import (
	"context"
	"depths/app/services"
	"depths/data_gateway"
)

type App struct {
	services    *services.Services
	dataGateway *data_gateway.DataGateway

	errorChan  chan error
	ctx        context.Context
	cancelFunc context.CancelFunc
}
