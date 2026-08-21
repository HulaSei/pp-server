package httpserver

import (
	"context"
	"crypto/tls"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/perfect-panel/server/internal/handler"
	"github.com/perfect-panel/server/internal/middleware"
	"github.com/perfect-panel/server/internal/module/notification"
	"github.com/perfect-panel/server/internal/route"
	"github.com/perfect-panel/server/pkg/logger"
)

type Server struct {
	h *server.Hertz
}

type Dependencies struct {
	Routes           route.Dependencies
	Notification     notification.Service
	TelegramBotToken func() string
}

func New(deps Dependencies, addr string, tlsConfig *tls.Config) *Server {
	opts := []config.Option{
		server.WithHostPorts(addr),
		server.WithDisablePrintRoute(true),
	}
	if tlsConfig != nil {
		opts = append(opts, server.WithTLS(tlsConfig))
	}

	return newServer(deps, opts)
}

func newServer(deps Dependencies, opts []config.Option) *Server {
	engine := server.Default(opts...)
	engine.Use(middleware.TraceMiddleware(), middleware.LoggerMiddleware(), middleware.CorsMiddleware)

	route.RegisterHandlers(engine, deps.Routes)
	handler.RegisterTelegramHandlers(engine, deps.Notification, deps.TelegramBotToken)
	handler.RegisterNotifyHandlers(engine, deps.Routes.Store, deps.Routes.Billing)

	return &Server{h: engine}
}

func (s *Server) Start() {
	if err := s.h.Run(); err != nil {
		logger.Errorf("server start error: %s", err.Error())
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.h.Shutdown(ctx)
}

func (s *Server) Engine() *server.Hertz {
	return s.h
}
