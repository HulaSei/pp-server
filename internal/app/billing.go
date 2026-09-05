package app

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/hibiken/asynq"
	"github.com/perfect-panel/server/internal/config"
	"github.com/perfect-panel/server/internal/infra/taskqueue"
	"github.com/perfect-panel/server/internal/module/billing"
	"github.com/perfect-panel/server/internal/repository"
	"github.com/redis/go-redis/v9"
)

// newBillingModule wires the billing module against the legacy store and the
// asynq client (ADR-001 step 4).
func newBillingModule(c config.Config, store repository.Store, queue *taskqueue.Client, rds *redis.Client, rate *billing.CurrencyRateCache, srv *Application) billing.Service {
	return billing.New(billing.Deps{
		Orders:       store.Order(),
		Payments:     store.Payment(),
		Coupons:      store.Coupon(),
		Withdrawals:  store.UserWithdrawal(),
		Plans:        store.Subscribe(),
		UserSubs:     store.UserSubscription(),
		Store:        store,
		Tx:           store,
		Queue:        activationQueue{client: queue},
		SingleModel:  func() bool { return srv.Runtime.Config().Subscribe.SingleModel },
		CurrencyUnit: func() string { return srv.Runtime.Config().Currency.Unit },
		Host:         c.Host,

		Logs:        store.Log(),
		UserCache:   store.UserCache(),
		Affiliates:  store.User(),
		AuthMethods: store.UserAuth(),

		UserProfiles: store.User(),
		InvitePolicy: func() (uint8, bool) {
			current := srv.Runtime.Config().Invite
			return uint8(current.ReferralPercentage), current.OnlyFirstPurchase
		},

		PortalPlans:        store.Subscribe(),
		GuestAccounts:      store.UserAuth(),
		Sessions:           rds,
		GuestCheckoutCache: rds,
		ActivationQueue:    queue,
		ExchangeRate:       rate,
		Portal: billing.PortalConfig{
			Host:              c.Host,
			SiteName:          func() string { return srv.Runtime.Config().Site.SiteName },
			CurrencyUnit:      func() string { return srv.Runtime.Config().Currency.Unit },
			CurrencyAccessKey: func() string { return srv.Runtime.Config().Currency.AccessKey },
			JwtSecret:         c.JwtAuth.AccessSecret,
			JwtExpire:         c.JwtAuth.AccessExpire,
		},
	})
}

// activationQueue adapts the asynq client to the billing module's activation
// port. A task-id conflict means a delivery already exists for the order,
// which is success, not an error.
type activationQueue struct {
	client *taskqueue.Client
}

func (q activationQueue) EnqueueActivation(ctx context.Context, orderNo string) error {
	payload, err := json.Marshal(taskqueue.ForthwithActivateOrderPayload{OrderNo: orderNo})
	if err != nil {
		return err
	}
	task := asynq.NewTask(taskqueue.ForthwithActivateOrder, payload)
	_, err = q.client.EnqueueContext(ctx, task, asynq.TaskID(taskqueue.ActivationTaskID(orderNo)))
	if errors.Is(err, asynq.ErrTaskIDConflict) {
		return nil
	}
	return err
}

// EnqueueDeferredClose schedules the pending order's expiry close after the
// payment window elapses.
func (q activationQueue) EnqueueDeferredClose(ctx context.Context, orderNo string) error {
	payload, err := json.Marshal(taskqueue.DeferCloseOrderPayload{OrderNo: orderNo})
	if err != nil {
		return err
	}
	task := asynq.NewTask(taskqueue.DeferCloseOrder, payload)
	_, err = q.client.EnqueueContext(ctx, task, asynq.MaxRetry(3), asynq.ProcessIn(billing.CloseOrderTimeMinutes*time.Minute))
	return err
}
