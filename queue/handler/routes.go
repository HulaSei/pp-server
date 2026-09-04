package handler

import (
	"github.com/hibiken/asynq"
	"github.com/perfect-panel/server/internal/eventbus"
	"github.com/perfect-panel/server/internal/module/billing"
	moduleSubscription "github.com/perfect-panel/server/internal/module/subscription"
	"github.com/perfect-panel/server/internal/repository"
	emailLogic "github.com/perfect-panel/server/queue/logic/email"
	"github.com/perfect-panel/server/queue/logic/events"
	orderLogic "github.com/perfect-panel/server/queue/logic/order"
	smslogic "github.com/perfect-panel/server/queue/logic/sms"
	subscriptionLogic "github.com/perfect-panel/server/queue/logic/subscription"
	"github.com/perfect-panel/server/queue/logic/task"
	"github.com/perfect-panel/server/queue/logic/traffic"
	"github.com/perfect-panel/server/queue/types"
)

type Dependencies struct {
	Email        emailLogic.Dependencies
	SMS          smslogic.Dependencies
	Order        orderLogic.Dependencies
	EventBus     *eventbus.Bus
	Traffic      traffic.Dependencies
	Subscription moduleSubscription.Service
	Store        repository.Store
	ExchangeRate *billing.CurrencyRateCache
}

func RegisterHandlers(mux *asynq.ServeMux, deps Dependencies) {
	var taskRepo repository.TaskRepo
	if deps.Store != nil {
		taskRepo = deps.Store.Task()
	}
	// Send email task
	mux.Handle(types.ForthwithSendEmail, emailLogic.NewSendEmailLogic(deps.Email))
	// Send sms task
	mux.Handle(types.ForthwithSendSms, smslogic.NewSendSmsLogic(deps.SMS))
	// Defer close order task
	mux.Handle(types.DeferCloseOrder, orderLogic.NewDeferCloseOrderLogic(deps.Order))
	// Forthwith activate order task
	mux.Handle(types.ForthwithActivateOrder, orderLogic.NewActivateOrderLogic(deps.Order))
	// Recover paid orders whose activation enqueue was interrupted.
	mux.Handle(types.SchedulerReconcilePaidOrders, orderLogic.NewReconcilePaidOrdersLogic(deps.Order))
	// Close stale pending orders even when their one-shot deferred task was
	// lost during a Redis outage or exhausted its retries.
	mux.Handle(types.SchedulerReconcilePendingOrders, orderLogic.NewReconcilePendingOrdersLogic(deps.Order))
	// Deliver durable order events to Redis Pub/Sub. The database remains the
	// source of truth for SSE replay when publication is delayed or duplicated.
	mux.Handle(types.SchedulerPublishOrderEvents, orderLogic.NewPublishOrderEventsLogic(deps.Order))
	// Domain events: the pump publishes outbox rows onto the queue; the
	// delivery worker runs the topic's subscribers per event.
	mux.Handle(types.SchedulerDispatchDomainEvents, events.NewDispatchDomainEventsLogic(deps.EventBus))
	mux.Handle(types.EventDeliver, events.NewDeliverDomainEventLogic(deps.EventBus))
	mux.Handle(types.SchedulerCleanupOrderEvents, orderLogic.NewCleanupOrderEventsLogic(deps.Order))
	// Daily settlement summary for administrators bound on Telegram.
	mux.Handle(types.SchedulerDailyOrderReport, orderLogic.NewDailyOrderReportLogic(deps.Order))

	// Forthwith traffic statistics
	mux.Handle(types.ForthwithTrafficStatistics, traffic.NewTrafficStatisticsLogic(deps.Traffic))
	// Flush aggregated traffic
	mux.Handle(types.SchedulerFlushTraffic, traffic.NewFlushTrafficLogic(deps.Traffic))

	// Schedule check subscription
	mux.Handle(types.SchedulerCheckSubscription, subscriptionLogic.NewCheckSubscriptionLogic(deps.Subscription))
	// Warn owners before their subscription expires.
	mux.Handle(types.SchedulerRemindExpiringSubscriptions, subscriptionLogic.NewRemindExpiringLogic(deps.Subscription))

	// Schedule total server data
	mux.Handle(types.SchedulerTotalServerData, traffic.NewServerDataLogic(deps.Traffic))

	// Schedule reset traffic
	mux.Handle(types.SchedulerResetTraffic, traffic.NewResetTrafficLogic(deps.Traffic))

	// ScheduledBatchSendEmail
	mux.Handle(types.ScheduledBatchSendEmail, emailLogic.NewBatchEmailLogic(deps.Email))

	// ScheduledTrafficStat
	mux.Handle(types.SchedulerTrafficStat, traffic.NewStatLogic(deps.Traffic))
	// ScheduledLogCleanup is independent from traffic aggregation so either
	// task can retry without suppressing the other.
	mux.Handle(types.SchedulerLogCleanup, traffic.NewLogCleanupLogic(deps.Traffic.Store, deps.Traffic.Log))

	// ForthwithQuotaTask
	mux.Handle(types.ForthwithQuotaTask, task.NewQuotaTaskLogic(deps.Subscription, taskRepo))
	// SchedulerExchangeRate
	mux.Handle(types.SchedulerExchangeRate, task.NewRateLogic(task.RateDependencies{Store: deps.Store, ExchangeRate: deps.ExchangeRate}))
}
