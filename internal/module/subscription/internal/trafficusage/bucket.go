// Package trafficusage owns idempotent subscription usage accounting.
package trafficusage

import (
	"context"

	"github.com/perfect-panel/server/internal/module/network/entity/traffic"
	"github.com/perfect-panel/server/internal/repository"
)

// BucketConsumer is persisted in the inbox; changing it replays old usage.
const BucketConsumer = "subscription.traffic_bucket"

type Store interface {
	Inbox() repository.InboxRepo
	InSubscriptionTx(context.Context, func(repository.SubscriptionStore) error) error
}

// ApplyBucketOnce commits usage and its inbox marker together. The network
// pipeline can retry its separate log transaction without charging usage again.
func ApplyBucketOnce(ctx context.Context, store Store, bucket string, deltas []traffic.SubscribeTrafficDelta) error {
	mark, err := store.Inbox().Find(ctx, BucketConsumer, bucket)
	if err != nil || mark != nil {
		return err
	}
	return store.InSubscriptionTx(ctx, func(tx repository.SubscriptionStore) error {
		if err := tx.UserSubscription().BatchUpdateUserSubscribeWithTraffic(ctx, deltas); err != nil {
			return err
		}
		return tx.Inbox().Insert(ctx, BucketConsumer, bucket, "")
	})
}
