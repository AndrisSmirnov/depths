package services

import (
	"depths/app/services/binance/service_binance_type"
	"depths/app/services/kuna/service_kuna_type"
)

type Services struct {
	BinanceService service_binance_type.IBinanceService
	KunaService    service_kuna_type.IKunaService
}
