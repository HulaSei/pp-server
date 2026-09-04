package initialize

import (
	"context"
	"errors"
	"time"

	"github.com/perfect-panel/server/initialize/migrate"
	"github.com/perfect-panel/server/internal/auth/password"
	"github.com/perfect-panel/server/internal/module/identity/entity/user"
	"github.com/perfect-panel/server/internal/repository"
	"github.com/perfect-panel/server/pkg/logger"
	"github.com/perfect-panel/server/pkg/orm"
)

func Migrate(ctx *Dependencies) {
	current := ctx.currentConfig()
	mc := orm.Mysql{
		Config: current.DatabaseConfig(),
	}
	now := time.Now()
	if err := migrate.Up(mc.Driver(), mc.MigrationDsn()); err != nil {
		if errors.Is(err, migrate.NoChange) {
			logger.Info("[Migrate] database not change")
			return
		}
		logger.Errorf("[Migrate] Up error: %v", err.Error())
		panic(err)
	} else {
		logger.Info("[Migrate] Database change, took " + time.Since(now).String())
	}
	// if not found admin user
	err := ctx.Store.InTx(context.Background(), func(store repository.Store) error {
		count, err := store.User().QueryResisterUserTotal(context.Background())
		if err != nil {
			return err
		}
		if count == 0 {
			enable := true
			admin := &user.User{
				Password:  password.EncodePassWord(current.Administrator.Password),
				Algo:      password.PasswordAlgoArgon2id,
				IsAdmin:   &enable,
				ReferCode: user.GenerateInviteCode(time.Now().Unix()),
			}
			if err := store.User().Insert(context.Background(), admin); err != nil {
				logger.Errorf("[Migrate] CreateAdminUser error: %v", err.Error())
				return err
			}
			if err := store.UserAuth().InsertUserAuthMethods(context.Background(), &user.AuthMethods{
				UserId:         admin.Id,
				AuthType:       "email",
				AuthIdentifier: current.Administrator.Email,
				Verified:       true,
			}); err != nil {
				logger.Errorf("[Migrate] CreateAdminUser error: %v", err.Error())
				return err
			}
			logger.Info("[Migrate] Create admin user success")
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
