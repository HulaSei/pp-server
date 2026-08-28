package traffic

import (
	"context"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
	"github.com/perfect-panel/server/internal/config"
	"github.com/perfect-panel/server/internal/repository"
	"github.com/perfect-panel/server/pkg/logger"
	"github.com/perfect-panel/server/pkg/timeutil"
)

type networkTransactor interface {
	InNetworkTx(ctx context.Context, fn func(repository.NetworkStore) error) error
}

// LogCleanupLogic owns retention independently from traffic aggregation so a
// failed statistics run cannot silently disable log cleanup (or vice versa).
type LogCleanupLogic struct {
	store networkTransactor
	log   func() config.Log
}

func NewLogCleanupLogic(store networkTransactor, logConfig func() config.Log) *LogCleanupLogic {
	return &LogCleanupLogic{store: store, log: logConfig}
}

func (l *LogCleanupLogic) ProcessTask(ctx context.Context, _ *asynq.Task) error {
	if l.store == nil || l.log == nil {
		return nil
	}
	settings := l.log()
	if !settings.AutoClear {
		return nil
	}
	if settings.ClearDays < 1 || settings.ClearDays > 3650 {
		return fmt.Errorf("invalid log retention days: %d", settings.ClearDays)
	}

	now := timeutil.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	threshold := today.AddDate(0, 0, -int(settings.ClearDays))
	if err := l.store.InNetworkTx(ctx, func(store repository.NetworkStore) error {
		if err := store.TrafficLog().DeleteBefore(ctx, threshold); err != nil {
			return err
		}
		return store.Log().DeleteBefore(ctx, threshold)
	}); err != nil {
		logger.WithContext(ctx).Errorw("[Log Cleanup] cleanup failed", logger.Field("error", err.Error()))
		return err
	}
	logger.WithContext(ctx).Infow("[Log Cleanup] cleanup completed", logger.Field("threshold", threshold.Format(time.DateOnly)))
	return nil
}
