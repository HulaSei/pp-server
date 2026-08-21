package traffic

import (
	"github.com/perfect-panel/server/internal/config"
	"github.com/perfect-panel/server/internal/repository"
	"github.com/perfect-panel/server/internal/trafficagg"
	"github.com/perfect-panel/server/pkg/asynqx"
	"github.com/redis/go-redis/v9"
)

type Dependencies struct {
	Store      repository.Store
	Redis      *redis.Client
	Queue      *asynqx.Client
	Log        func() config.Log
	Aggregator trafficagg.Deps
}
