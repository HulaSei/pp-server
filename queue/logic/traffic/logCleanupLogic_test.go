package traffic

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/perfect-panel/server/internal/config"
	"github.com/perfect-panel/server/internal/repository"
)

type cleanupTrafficRepo struct {
	repository.TrafficRepo
	threshold time.Time
	err       error
}

func (r *cleanupTrafficRepo) DeleteBefore(_ context.Context, threshold time.Time) error {
	r.threshold = threshold
	return r.err
}

type cleanupLogRepo struct {
	repository.LogRepo
	threshold time.Time
	err       error
}

func (r *cleanupLogRepo) DeleteBefore(_ context.Context, threshold time.Time) error {
	r.threshold = threshold
	return r.err
}

type cleanupNetworkStore struct {
	repository.NetworkStore
	traffic repository.TrafficRepo
	logs    repository.LogRepo
}

func (s cleanupNetworkStore) TrafficLog() repository.TrafficRepo { return s.traffic }
func (s cleanupNetworkStore) Log() repository.LogRepo            { return s.logs }

type cleanupTransactor struct {
	store repository.NetworkStore
	calls int
}

func (t *cleanupTransactor) InNetworkTx(ctx context.Context, fn func(repository.NetworkStore) error) error {
	t.calls++
	return fn(t.store)
}

func TestLogCleanupRunsIndependentlyAndPropagatesFailures(t *testing.T) {
	trafficRepo := &cleanupTrafficRepo{}
	logRepo := &cleanupLogRepo{}
	tx := &cleanupTransactor{store: cleanupNetworkStore{traffic: trafficRepo, logs: logRepo}}
	logic := NewLogCleanupLogic(tx, func() config.Log { return config.Log{AutoClear: true, ClearDays: 7} })

	if err := logic.ProcessTask(context.Background(), nil); err != nil {
		t.Fatal(err)
	}
	if tx.calls != 1 || trafficRepo.threshold.IsZero() || !trafficRepo.threshold.Equal(logRepo.threshold) {
		t.Fatalf("cleanup was not applied atomically: calls=%d traffic=%v logs=%v", tx.calls, trafficRepo.threshold, logRepo.threshold)
	}

	trafficRepo.err = errors.New("delete failed")
	if err := logic.ProcessTask(context.Background(), nil); !errors.Is(err, trafficRepo.err) {
		t.Fatalf("cleanup error = %v, want %v", err, trafficRepo.err)
	}
}

func TestLogCleanupRejectsUnsafeRetentionWithoutDeleting(t *testing.T) {
	tx := &cleanupTransactor{store: cleanupNetworkStore{}}
	logic := NewLogCleanupLogic(tx, func() config.Log { return config.Log{AutoClear: true, ClearDays: -1} })
	if err := logic.ProcessTask(context.Background(), nil); err == nil {
		t.Fatal("unsafe retention was accepted")
	}
	if tx.calls != 0 {
		t.Fatalf("cleanup transaction ran %d times", tx.calls)
	}
}
