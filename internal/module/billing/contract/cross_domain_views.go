package dto

// Cross-domain view types are module-owned snapshots. They deliberately
// duplicate JSON shapes instead of coupling one module's API to another.
import (
	bytes "bytes"
	json "encoding/json"
	strconv "strconv"
	strings "strings"
)

type BillingBalanceLogSnapshot struct {
	Type      uint16 `json:"type"`
	UserId    int64  `json:"user_id"`
	Amount    int64  `json:"amount"`
	OrderNo   string `json:"order_no,omitempty"`
	Balance   int64  `json:"balance"`
	Timestamp int64  `json:"timestamp"`
} // @name dto.BalanceLog

type BillingCommissionLogSnapshot struct {
	Type      uint16 `json:"type"`
	UserId    int64  `json:"user_id"`
	Amount    int64  `json:"amount"`
	OrderNo   string `json:"order_no"`
	Timestamp int64  `json:"timestamp"`
} // @name dto.CommissionLog

type CommissionWithdrawRequest struct {
	Amount  int64  `json:"amount" validate:"required,gt=0,lte=2000000000"`
	Content string `json:"content"`
}

type GetSubscriptionRequest struct {
	Language string `form:"language"`
}

type GetSubscriptionResponse struct {
	List []BillingSubscribeSnapshot `json:"list"`
}

type QueryUserAffiliateCountResponse struct {
	Registers       int64 `json:"registers"`
	TotalCommission int64 `json:"total_commission"`
}

type QueryUserAffiliateListRequest struct {
	Page int `form:"page" validate:"required,gt=0"`
	Size int `form:"size" validate:"required,gt=0,lte=100"`
}

type QueryUserAffiliateListResponse struct {
	List  []UserAffiliate `json:"list"`
	Total int64           `json:"total"`
}

type QueryUserBalanceLogListResponse struct {
	List  []BillingBalanceLogSnapshot `json:"list"`
	Total int64                       `json:"total"`
}

type QueryUserCommissionLogListRequest struct {
	Page int `form:"page" validate:"required,gt=0"`
	Size int `form:"size" validate:"required,gt=0,lte=100"`
}

type QueryUserCommissionLogListResponse struct {
	List  []BillingCommissionLogSnapshot `json:"list"`
	Total int64                          `json:"total"`
}

type QueryWithdrawalLogListRequest struct {
	Page int `form:"page" validate:"required,gt=0"`
	Size int `form:"size" validate:"required,gt=0,lte=100"`
}

type QueryWithdrawalLogListResponse struct {
	List  []WithdrawalLog `json:"list"`
	Total int64           `json:"total"`
}

type BillingNodeIDList []int64 // @name dto.StringInt64Slice

func (s *BillingNodeIDList) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*s = nil
		return nil
	}

	var raw []json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	values := make([]int64, 0, len(raw))
	for _, item := range raw {
		var number int64
		if err := json.Unmarshal(item, &number); err == nil {
			values = append(values, number)
			continue
		}

		var text string
		if err := json.Unmarshal(item, &text); err != nil {
			return err
		}
		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}
		parsed, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			return err
		}
		values = append(values, parsed)
	}

	*s = values
	return nil
}

func (s BillingNodeIDList) MarshalJSON() ([]byte, error) {
	values := make([]string, 0, len(s))
	for _, item := range s {
		values = append(values, strconv.FormatInt(item, 10))
	}
	return json.Marshal(values)
}

func (s BillingNodeIDList) Int64s() []int64 {
	return []int64(s)
}

type BillingSubscribeSnapshot struct {
	Id                int64                      `json:"id"`
	Name              string                     `json:"name"`
	Language          string                     `json:"language"`
	Description       string                     `json:"description"`
	UnitPrice         int64                      `json:"unit_price"`
	UnitTime          string                     `json:"unit_time"`
	Discount          []BillingSubscribeDiscount `json:"discount"`
	Replacement       int64                      `json:"replacement"`
	Inventory         int64                      `json:"inventory"`
	Traffic           int64                      `json:"traffic"`
	SpeedLimit        int64                      `json:"speed_limit"`
	DeviceLimit       int64                      `json:"device_limit"`
	Quota             int64                      `json:"quota"`
	Nodes             BillingNodeIDList          `json:"nodes"`
	NodeTags          []string                   `json:"node_tags"`
	Show              bool                       `json:"show"`
	Sell              bool                       `json:"sell"`
	Sort              int64                      `json:"sort"`
	DeductionRatio    int64                      `json:"deduction_ratio"`
	AllowDeduction    bool                       `json:"allow_deduction"`
	ResetCycle        int64                      `json:"reset_cycle"`
	RenewalReset      bool                       `json:"renewal_reset"`
	ShowOriginalPrice bool                       `json:"show_original_price"`
	CreatedAt         int64                      `json:"created_at"`
	UpdatedAt         int64                      `json:"updated_at"`
} // @name dto.Subscribe

type UserAffiliate struct {
	Avatar       string `json:"avatar"`
	Identifier   string `json:"identifier"`
	RegisteredAt int64  `json:"registered_at"`
	Enable       bool   `json:"enable"`
}

type WithdrawalLog struct {
	Id        int64  `json:"id"`
	UserId    int64  `json:"user_id"`
	Amount    int64  `json:"amount"`
	Content   string `json:"content"`
	Status    uint8  `json:"status"`
	Reason    string `json:"reason,omitempty"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
