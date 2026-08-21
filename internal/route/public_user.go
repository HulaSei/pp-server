package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	publicUser "github.com/perfect-panel/server/internal/handler/public/user"
)

func registerPublicUserRoutes(router *server.Hertz, deps Dependencies) {
	publicUserGroupRouter := router.Group("/v1/public/user")
	publicUserGroupRouter.Use(deps.authMiddleware(), deps.deviceMiddleware())
	publicUserGroupRouter.GET("/affiliate/count", publicUser.QueryUserAffiliateHandler(deps.Billing))
	publicUserGroupRouter.GET("/affiliate/list", publicUser.QueryUserAffiliateListHandler(deps.Billing))
	publicUserGroupRouter.GET("/balance_log", publicUser.QueryUserBalanceLogHandler(deps.Billing))
	publicUserGroupRouter.PUT("/bind_email", publicUser.UpdateBindEmailHandler(deps.Identity))
	publicUserGroupRouter.PUT("/bind_mobile", publicUser.UpdateBindMobileHandler(deps.Identity))
	publicUserGroupRouter.POST("/bind_oauth", publicUser.BindOAuthHandler(deps.Identity))
	publicUserGroupRouter.POST("/bind_oauth/callback", publicUser.BindOAuthCallbackHandler(deps.Identity))
	publicUserGroupRouter.GET("/bind_telegram", publicUser.BindTelegramHandler(deps.Identity))
	publicUserGroupRouter.GET("/commission_log", publicUser.QueryUserCommissionLogHandler(deps.Billing))
	publicUserGroupRouter.POST("/commission_withdraw", publicUser.CommissionWithdrawHandler(deps.Billing))
	publicUserGroupRouter.GET("/devices", publicUser.GetDeviceListHandler(deps.Identity))
	publicUserGroupRouter.GET("/info", publicUser.QueryUserInfoHandler(deps.Identity))
	publicUserGroupRouter.GET("/login_log", publicUser.GetLoginLogHandler(deps.Identity))
	publicUserGroupRouter.PUT("/notify", publicUser.UpdateUserNotifyHandler(deps.Identity))
	publicUserGroupRouter.GET("/oauth_methods", publicUser.GetOAuthMethodsHandler(deps.Identity))
	publicUserGroupRouter.PUT("/password", publicUser.UpdateUserPasswordHandler(deps.Identity))
	publicUserGroupRouter.PUT("/rules", publicUser.UpdateUserRulesHandler(deps.Identity))
	publicUserGroupRouter.GET("/subscribe", publicUser.QueryUserSubscribeHandler(deps.Subscription))
	publicUserGroupRouter.GET("/subscribe_log", publicUser.GetSubscribeLogHandler(deps.Subscription))
	publicUserGroupRouter.PUT("/subscribe_note", publicUser.UpdateUserSubscribeNoteHandler(deps.Subscription))
	publicUserGroupRouter.PUT("/subscribe_token", publicUser.ResetUserSubscribeTokenHandler(deps.Subscription))
	publicUserGroupRouter.PUT("/unbind_device", publicUser.UnbindDeviceHandler(deps.Identity))
	publicUserGroupRouter.POST("/unbind_oauth", publicUser.UnbindOAuthHandler(deps.Identity))
	publicUserGroupRouter.POST("/unbind_telegram", publicUser.UnbindTelegramHandler(deps.Identity))
	publicUserGroupRouter.POST("/unsubscribe", publicUser.UnsubscribeHandler(deps.Subscription))
	publicUserGroupRouter.POST("/unsubscribe/pre", publicUser.PreUnsubscribeHandler(deps.Subscription))
	publicUserGroupRouter.POST("/verify_email", publicUser.VerifyEmailHandler(deps.Identity))
	publicUserGroupRouter.GET("/withdrawal_log", publicUser.QueryWithdrawalLogHandler(deps.Billing))
}
