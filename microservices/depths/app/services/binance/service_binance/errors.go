package service_binance

import (
	"depths/pkg/errors"
	"depths/voc"
	"fmt"
)

var (
	ErrNilExchangeManager = errors.NewCriticalErrorWithMessage(
		errors.ErrDBUnknown,
		fmt.Sprintf("%v: %v", voc.BinanceService, voc.NilExchangeManager))
)
