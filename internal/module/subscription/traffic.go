package subscription

import (
	"context"

	"github.com/perfect-panel/server/internal/module/network/entity/traffic"
	"github.com/perfect-panel/server/internal/module/subscription/internal/trafficusage"
)

type TrafficUsageStore = trafficusage.Store

func ApplyTrafficBucketOnce(ctx context.Context, store TrafficUsageStore, bucket string, deltas []traffic.SubscribeTrafficDelta) error {
	return trafficusage.ApplyBucketOnce(ctx, store, bucket, deltas)
}
