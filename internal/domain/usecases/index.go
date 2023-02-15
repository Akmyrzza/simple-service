package usecases

import (
	"github.com/Akmyrzza/simple-service/internal/adapters/logger"
	"github.com/Akmyrzza/simple-service/internal/domain/core"
)

type St struct {
	lg logger.WarnAndError
	cr *core.St
}

func New(lg logger.WarnAndError, cr *core.St) *St {
	return &St{
		lg: lg,
		cr: cr,
	}
}
