package service_kuna

import (
	"depths/pkg/errors"
	"depths/voc"
	"fmt"
)

var (
	ErrNilExchangeManager = errors.NewCriticalErrorWithMessage(
		errors.ErrDBUnknown,
		fmt.Sprintf("%v: %v", voc.KunaService, voc.NilExchangeManager))
)
