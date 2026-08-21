package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	adminUser "github.com/perfect-panel/server/internal/handler/admin/user"
)

func registerAdminUserRoutes(router *server.Hertz, deps Dependencies) {
	adminUserGroupRouter := router.Group("/v1/admin/user")
	adminUserGroupRouter.Use(deps.authMiddleware())
	{
		adminUserGroupRouter.DELETE("/", adminUser.DeleteUserHandler(deps.Identity))
		adminUserGroupRouter.POST("/", adminUser.CreateUserHandler(deps.Identity))
		adminUserGroupRouter.POST("/auth_method", adminUser.CreateUserAuthMethodHandler(deps.Identity))
		adminUserGroupRouter.DELETE("/auth_method", adminUser.DeleteUserAuthMethodHandler(deps.Identity))
		adminUserGroupRouter.PUT("/auth_method", adminUser.UpdateUserAuthMethodHandler(deps.Identity))
		adminUserGroupRouter.GET("/auth_method", adminUser.GetUserAuthMethodHandler(deps.Identity))
		adminUserGroupRouter.PUT("/basic", adminUser.UpdateUserBasicInfoHandler(deps.Identity))
		adminUserGroupRouter.DELETE("/batch", adminUser.BatchDeleteUserHandler(deps.Identity))
		adminUserGroupRouter.GET("/current", adminUser.CurrentUserHandler(deps.Identity))
		adminUserGroupRouter.GET("/detail", adminUser.GetUserDetailHandler(deps.Identity))
		adminUserGroupRouter.PUT("/device", adminUser.UpdateUserDeviceHandler(deps.Identity))
		adminUserGroupRouter.DELETE("/device", adminUser.DeleteUserDeviceHandler(deps.Identity))
		adminUserGroupRouter.PUT("/device/kick_offline", adminUser.KickOfflineByUserDeviceHandler(deps.Identity))
		adminUserGroupRouter.GET("/list", adminUser.GetUserListHandler(deps.Identity))
		adminUserGroupRouter.GET("/login/logs", adminUser.GetUserLoginLogsHandler(deps.Identity))
		adminUserGroupRouter.PUT("/notify", adminUser.UpdateUserNotifySettingHandler(deps.Identity))
		adminUserGroupRouter.GET("/subscribe", adminUser.GetUserSubscribeHandler(deps.Subscription))
		adminUserGroupRouter.POST("/subscribe", adminUser.CreateUserSubscribeHandler(deps.Subscription))
		adminUserGroupRouter.PUT("/subscribe", adminUser.UpdateUserSubscribeHandler(deps.Subscription))
		adminUserGroupRouter.DELETE("/subscribe", adminUser.DeleteUserSubscribeHandler(deps.Subscription))
		adminUserGroupRouter.GET("/subscribe/detail", adminUser.GetUserSubscribeByIdHandler(deps.Subscription))
		adminUserGroupRouter.GET("/subscribe/device", adminUser.GetUserSubscribeDevicesHandler(deps.Subscription))
		adminUserGroupRouter.GET("/subscribe/logs", adminUser.GetUserSubscribeLogsHandler(deps.Subscription))
		adminUserGroupRouter.GET("/subscribe/reset/logs", adminUser.GetUserSubscribeResetTrafficLogsHandler(deps.Subscription))
		adminUserGroupRouter.POST("/subscribe/reset/token", adminUser.ResetUserSubscribeTokenHandler(deps.Subscription))
		adminUserGroupRouter.POST("/subscribe/reset/traffic", adminUser.ResetUserSubscribeTrafficHandler(deps.Subscription))
		adminUserGroupRouter.POST("/subscribe/toggle", adminUser.ToggleUserSubscribeStatusHandler(deps.Subscription))
		adminUserGroupRouter.GET("/subscribe/traffic_logs", adminUser.GetUserSubscribeTrafficLogsHandler(deps.Subscription))
	}
}
