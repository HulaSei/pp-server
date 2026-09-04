package traffic

import (
	"github.com/perfect-panel/server/internal/config"
	"github.com/perfect-panel/server/internal/repository"
	"github.com/perfect-panel/server/internal/taskqueue"
	"github.com/perfect-panel/server/internal/trafficagg"
	"github.com/redis/go-redis/v9"
)

type Dependencies struct {
	Store      repository.Store
	Redis      *redis.Client
	Queue      *taskqueue.Client
	Log        func() config.Log
	Aggregator trafficagg.Deps
}
