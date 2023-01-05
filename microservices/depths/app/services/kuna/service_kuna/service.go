package service_kuna

import (
	"depths/app/domain/exchange_manager_domain"
	"depths/app/services/kuna/service_kuna_type"
)

type kunaService struct {
	name    string
	manager exchange_manager_domain.IExchangeManagerSub
}

func (k *kunaService) GetName() string {
	return k.name
}

func NewKunaService(
	manager exchange_manager_domain.IExchangeManagerSub,
) (
	service_kuna_type.IKunaService, error,
) {
	if manager == nil {
		return nil, ErrNilExchangeManager
	}

	k := &kunaService{
		name:    "Kuna",
		manager: manager,
	}

	k.manager.RegisterManager(k)

	return k, nil
}
