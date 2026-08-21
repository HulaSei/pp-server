package dto

import "github.com/perfect-panel/server/pkg/platform"

// Cross-domain view types are module-owned snapshots. They deliberately
// duplicate JSON shapes instead of coupling one module's API to another.
type GetLoginLogRequest struct {
	Page int `form:"page" validate:"required,gt=0"`
	Size int `form:"size" validate:"required,gt=0,lte=100"`
}

type GetLoginLogResponse struct {
	List  []UserLoginLog `json:"list"`
	Total int64          `json:"total"`
}

type GetUserLoginLogsRequest struct {
	Page   int   `form:"page" validate:"required,gt=0"`
	Size   int   `form:"size" validate:"required,gt=0,lte=100"`
	UserId int64 `form:"user_id"`
}

type GetUserLoginLogsResponse struct {
	List  []UserLoginLog `json:"list"`
	Total int64          `json:"total"`
}

type AuthPlatformInfoSnapshot = platform.Info // @name dto.PlatformInfo

type AuthPlatformResponse struct {
	List []AuthPlatformInfoSnapshot `json:"list"`
} // @name dto.PlatformResponse

type UserLoginLog struct {
	Id        int64  `json:"id"`
	UserId    int64  `json:"user_id"`
	LoginIP   string `json:"login_ip"`
	UserAgent string `json:"user_agent"`
	Success   bool   `json:"success"`
	Timestamp int64  `json:"timestamp"`
}
