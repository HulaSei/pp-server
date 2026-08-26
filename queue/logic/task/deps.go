package task

import (
	"github.com/perfect-panel/server/internal/repository"
	"github.com/perfect-panel/server/pkg/exchangeRate"
)

type RateDependencies struct {
	Store        repository.Store
	ExchangeRate *exchangeRate.Cache
}
