package exchange_manager_domain

import (
	"depths/pkg/errors"
	"depths/voc"
	"fmt"
)

var (
	errNilDataGateway = errors.NewCriticalErrorWithMessage(
		errors.ErrDBUnknown,
		fmt.Sprintf("%v: %v", voc.ExchangeManager, voc.NilDataGateway))
)
