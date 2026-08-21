package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	adminCoupon "github.com/perfect-panel/server/internal/handler/admin/coupon"
	adminOrder "github.com/perfect-panel/server/internal/handler/admin/order"
	adminPayment "github.com/perfect-panel/server/internal/handler/admin/payment"
	adminSubscribe "github.com/perfect-panel/server/internal/handler/admin/subscribe"
)

func registerAdminCouponRoutes(router *server.Hertz, deps Dependencies) {
	group := router.Group("/v1/admin/coupon")
	group.Use(deps.authMiddleware())
	group.POST("/", adminCoupon.CreateCouponHandler(deps.Billing))
	group.PUT("/", adminCoupon.UpdateCouponHandler(deps.Billing))
	group.DELETE("/", adminCoupon.DeleteCouponHandler(deps.Billing))
	group.DELETE("/batch", adminCoupon.BatchDeleteCouponHandler(deps.Billing))
	group.GET("/list", adminCoupon.GetCouponListHandler(deps.Billing))
}

func registerAdminOrderRoutes(router *server.Hertz, deps Dependencies) {
	group := router.Group("/v1/admin/order")
	group.Use(deps.authMiddleware())
	group.POST("/", adminOrder.CreateOrderHandler(deps.Billing))
	group.GET("/list", adminOrder.GetOrderListHandler(deps.Billing))
	group.PUT("/status", adminOrder.UpdateOrderStatusHandler(deps.Billing))
}

func registerAdminPaymentRoutes(router *server.Hertz, deps Dependencies) {
	group := router.Group("/v1/admin/payment")
	group.Use(deps.authMiddleware())
	group.POST("/", adminPayment.CreatePaymentMethodHandler(deps.Billing))
	group.PUT("/", adminPayment.UpdatePaymentMethodHandler(deps.Billing))
	group.DELETE("/", adminPayment.DeletePaymentMethodHandler(deps.Billing))
	group.GET("/list", adminPayment.GetPaymentMethodListHandler(deps.Billing))
	group.GET("/platform", adminPayment.GetPaymentPlatformHandler(deps.Billing))
}

func registerAdminSubscribeRoutes(router *server.Hertz, deps Dependencies) {
	group := router.Group("/v1/admin/subscribe")
	group.Use(deps.authMiddleware())
	group.POST("/", adminSubscribe.CreateSubscribeHandler(deps.Subscription))
	group.PUT("/", adminSubscribe.UpdateSubscribeHandler(deps.Subscription))
	group.DELETE("/", adminSubscribe.DeleteSubscribeHandler(deps.Subscription))
	group.DELETE("/batch", adminSubscribe.BatchDeleteSubscribeHandler(deps.Subscription))
	group.GET("/details", adminSubscribe.GetSubscribeDetailsHandler(deps.Subscription))
	group.POST("/group", adminSubscribe.CreateSubscribeGroupHandler(deps.Subscription))
	group.PUT("/group", adminSubscribe.UpdateSubscribeGroupHandler(deps.Subscription))
	group.DELETE("/group", adminSubscribe.DeleteSubscribeGroupHandler(deps.Subscription))
	group.DELETE("/group/batch", adminSubscribe.BatchDeleteSubscribeGroupHandler(deps.Subscription))
	group.GET("/group/list", adminSubscribe.GetSubscribeGroupListHandler(deps.Subscription))
	group.GET("/list", adminSubscribe.GetSubscribeListHandler(deps.Subscription))
	group.POST("/reset_all_token", adminSubscribe.ResetAllSubscribeTokenHandler(deps.Subscription))
	group.POST("/sort", adminSubscribe.SubscribeSortHandler(deps.Subscription))
}
