package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/perfect-panel/server/internal/handler"
)

func registerSubscribeConfigRoutes(router *server.Hertz, deps Dependencies) {
	subscribePath := deps.runtimeConfig().Subscribe.SubscribePath
	if subscribePath == "" {
		subscribePath = "/v1/subscribe/config"
	}
	subscribeDeps := handler.SubscribeDeps{Service: deps.Subscription, Config: deps.subscribeConfig}
	router.GET(subscribePath, handler.SubscribeHandler(subscribeDeps))
	if deps.runtimeConfig().Subscribe.PanDomain {
		router.GET("/", handler.PanDomainSubscribeHandler(subscribeDeps))
	}
}
