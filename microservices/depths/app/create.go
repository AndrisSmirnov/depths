package app

import "context"

func CreateApplication(ctx context.Context, cancelFunc func()) (*App, error) {
	const ERROR_AMOUNT = 5

	errChan := make(chan error, ERROR_AMOUNT)

	dataGateway, err := CreateDataGateway(ctx, errChan)
	if err != nil {
		return nil, err
	}

	return &App{
		dataGateway: dataGateway,
		cancelFunc:  cancelFunc,
	}, nil
}
