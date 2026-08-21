package emailLogic

import (
	"github.com/perfect-panel/server/internal/config"
	"github.com/perfect-panel/server/internal/repository"
)

type Dependencies struct {
	Store    repository.Store
	Email    func() config.EmailConfig
	SiteName func() string
}
