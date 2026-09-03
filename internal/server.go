package internal

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"time"

	"github.com/perfect-panel/server/initialize"
	"github.com/perfect-panel/server/internal/config"
	"github.com/perfect-panel/server/internal/repository"
	"github.com/perfect-panel/server/internal/transport/httpserver"
	"github.com/perfect-panel/server/pkg/logger"

	"github.com/perfect-panel/server/pkg/proc"
	"github.com/perfect-panel/server/pkg/trace"
)

type Service struct {
	server transportServer
	deps   Dependencies
}

type Dependencies struct {
	Config                 func() config.Config
	Store                  repository.Store
	Initialize             *initialize.Dependencies
	HTTP                   func() httpserver.Dependencies
	SetRestart             func(func() error)
	SetReinitializeHandler func(func(string))
}

func NewService(deps Dependencies) *Service {
	return &Service{deps: deps}
}

type transportServer interface {
	Start()
	Shutdown(ctx context.Context) error
}

func newTransportServer(deps httpserver.Dependencies, runtimeConfig config.Config, addr string) transportServer {
	var tlsConfig *tls.Config
	if runtimeConfig.TLS.Enable {
		cert, err := tls.LoadX509KeyPair(runtimeConfig.TLS.CertFile, runtimeConfig.TLS.KeyFile)
		if err != nil {
			logger.Errorf("load tls certificate error: %s", err.Error())
			return nil
		}
		tlsConfig = &tls.Config{
			MinVersion:   tls.VersionTLS12,
			Certificates: []tls.Certificate{cert},
		}
	}
	return httpserver.New(deps, addr, tlsConfig)
}

func (m *Service) Start() {
	if m.deps.Config == nil || m.deps.Initialize == nil || m.deps.HTTP == nil {
		panic("config file path is nil")
	}

	runtimeConfig := m.deps.Config()
	serverAddr := fmt.Sprintf("%v:%d", runtimeConfig.Host, runtimeConfig.Port)
	initialize.StartInitSystemConfig(m.deps.Initialize)
	if err := m.deps.Store.UserAuth().ValidateEmailIdentityUniqueness(context.Background()); err != nil {
		panic(err.Error())
	}
	m.server = newTransportServer(m.deps.HTTP(), m.deps.Config(), serverAddr)
	if m.server == nil {
		return
	}
	traceConfig := runtimeConfig.Trace
	if traceConfig.Name == "" {
		traceConfig.Name = trace.TraceName
	}
	trace.StartAgent(traceConfig)
	proc.AddShutdownListener(func() {
		trace.StopAgent()
	})
	if m.deps.SetRestart != nil {
		m.deps.SetRestart(m.Restart)
	}
	reinitialize := func(subsystem string) {
		switch subsystem {
		case "verify":
			initialize.Verify(m.deps.Initialize)
		case "node":
			initialize.Node(m.deps.Initialize)
		case "telegram":
			initialize.Telegram(m.deps.Initialize)
		case "currency":
			initialize.Currency(m.deps.Initialize)
		case "register":
			initialize.Register(m.deps.Initialize)
		case "site":
			initialize.Site(m.deps.Initialize)
		case "invite":
			initialize.Invite(m.deps.Initialize)
		case "subscribe":
			initialize.Subscribe(m.deps.Initialize)
		case "email":
			initialize.Email(m.deps.Initialize)
		case "mobile":
			initialize.Mobile(m.deps.Initialize)
		case "device":
			initialize.Device(m.deps.Initialize)
		}
	}
	if m.deps.SetReinitializeHandler != nil {
		m.deps.SetReinitializeHandler(reinitialize)
	}
	logger.Infof("server start at %v", serverAddr)
	m.server.Start()
}

func (m *Service) Stop() {
	if m.server == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := m.server.Shutdown(ctx); err != nil {
		logger.Errorf("server shutdown error: %s", err.Error())
	}
	logger.Info("server shutdown")
}

func (m *Service) Restart() error {
	if m.server == nil {
		return errors.New("server is nil")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := m.server.Shutdown(ctx); err != nil {
		logger.Errorf("server shutdown error: %v", err.Error())
		return err
	}
	logger.Info("server shutdown")
	go m.Start()
	return nil
}
