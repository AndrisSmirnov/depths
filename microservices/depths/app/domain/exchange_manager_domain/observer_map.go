package exchange_manager_domain

import (
	"depths/app/domain/exchange_domain"
)

func (o ObserverMap) find(name string) exchange_domain.Observer {
	return o[name]
}
