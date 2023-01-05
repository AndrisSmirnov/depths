package app

import (
	"depths/app/domain/market_precision_domain"
	"encoding/json"
	"os"
)

func (a *App) Launch() error {
	if err := a.InitMarketPrecisions(); err != nil {
		return err
	}

	return a.depthManager.Start()
}

func (a *App) InitMarketPrecisions() error {
	mPrecisions, err := readFromFileMarketPrecisions()
	if err != nil {
		return err
	}

	return a.dataGateway.InsertManyMarketPrecisions(mPrecisions)
}

//goland:noinspection GoDeprecation
func readFromFileMarketPrecisions() ([]market_precision_domain.MarketPrecision, error) {
	var data []market_precision_domain.MarketPrecision

	file, err := os.ReadFile("marketprecisions.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
