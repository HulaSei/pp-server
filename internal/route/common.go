package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	common "github.com/perfect-panel/server/internal/handler/common"
)

func registerCommonRoutes(router *server.Hertz, deps Dependencies) {
	commonGroupRouter := router.Group("/v1/common")
	commonGroupRouter.Use(deps.deviceMiddleware())
	{
		commonGroupRouter.GET("/ads", common.GetAdsHandler(deps.Support))
		commonGroupRouter.POST("/check_verification_code", common.CheckVerificationCodeHandler(deps.Identity))
		commonGroupRouter.GET("/client", common.GetClientHandler(deps.Platform))
		commonGroupRouter.GET("/heartbeat", common.HeartbeatHandler(deps.Platform))
		commonGroupRouter.POST("/send_code", common.SendEmailCodeHandler(deps.Identity))
		commonGroupRouter.POST("/send_sms_code", common.SendSmsCodeHandler(deps.Identity))
		commonGroupRouter.GET("/site/config", common.GetGlobalConfigHandler(deps.Platform))
		commonGroupRouter.GET("/site/privacy", common.GetPrivacyPolicyHandler(deps.Platform))
		commonGroupRouter.GET("/site/stat", common.GetStatHandler(deps.Platform))
		commonGroupRouter.GET("/site/tos", common.GetTosHandler(deps.Platform))
	}
}
