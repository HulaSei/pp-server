package initialize

import (
	tgbot "github.com/go-telegram/bot"
	"github.com/perfect-panel/server/internal/config"
	"github.com/perfect-panel/server/internal/module/billing"
	"github.com/perfect-panel/server/internal/module/network"
	"github.com/perfect-panel/server/internal/module/notification"
	"github.com/perfect-panel/server/internal/repository"
)

// Dependencies is the startup/reconfiguration boundary. It owns only mutable
// runtime configuration and the services needed to load or publish it.
type Dependencies struct {
	Config                   func() config.Config
	UpdateConfig             func(func(*config.Config))
	Store                    repository.Store
	ExchangeRate             *billing.CurrencyRateCache
	Notification             notification.Service
	SetTelegramBot           func(*tgbot.Bot)
	SetNodeMultiplierManager func(*network.MultiplierManager)
}

func (d *Dependencies) currentConfig() config.Config {
	if d == nil || d.Config == nil {
		return config.Config{}
	}
	return d.Config()
}

func (d *Dependencies) updateConfig(update func(*config.Config)) {
	if d != nil && d.UpdateConfig != nil {
		d.UpdateConfig(update)
	}
}
