package smslogic

import (
	"context"
	"encoding/json"

	"github.com/perfect-panel/server/pkg/logger"
	"github.com/perfect-panel/server/pkg/timeutil"

	"github.com/hibiken/asynq"
	"github.com/perfect-panel/server/internal/module/platform/entity/log"
	"github.com/perfect-panel/server/pkg/constant"
	"github.com/perfect-panel/server/pkg/sms"
	"github.com/perfect-panel/server/queue/types"
)

type SmsSendCount struct {
	Count    int   `json:"count"`
	CreateAt int64 `json:"create_at"`
}

type SendSmsLogic struct {
	deps Dependencies
}

func newSMSMessageLog(platform string, messageType uint8) *log.Message {
	return &log.Message{
		Platform: platform,
		To:       logger.RedactedValue,
		Subject:  constant.ParseVerifyType(messageType).String(),
		Content:  map[string]interface{}{"redacted": true},
	}
}

func NewSendSmsLogic(deps Dependencies) *SendSmsLogic {
	return &SendSmsLogic{
		deps: deps,
	}
}
func (l *SendSmsLogic) ProcessTask(ctx context.Context, task *asynq.Task) error {
	var payload types.SendSmsPayload
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		logger.WithContext(ctx).Error("[SendSmsLogic] Unmarshal payload failed",
			logger.Field("error", err.Error()),
			logger.Field("payload", task.Payload()),
		)
		return nil
	}
	client, err := sms.NewSender(l.deps.Mobile().Platform, l.deps.Mobile().PlatformConfig)
	if err != nil {
		logger.WithContext(ctx).Error("[SendSmsLogic] New send sms client failed", logger.Field("error", err.Error()), logger.Field("payload", payload))
		return err
	}
	createSms := newSMSMessageLog(l.deps.Mobile().Platform, payload.Type)
	content, marshalErr := createSms.Marshal()
	if marshalErr != nil {
		return marshalErr
	}
	audit := &log.SystemLog{
		Type:     log.TypeMobileMessage.Uint8(),
		Date:     timeutil.Now().Format("2006-01-02"),
		ObjectID: 0,
		Content:  string(content),
	}
	// Record the attempt before contacting the provider so a storage failure
	// cannot produce a successful but unaudited SMS delivery.
	if err = l.deps.Store.Log().Insert(ctx, audit); err != nil {
		logger.WithContext(ctx).Error("[SendSmsLogic] Insert sms log failed", logger.Field("error", err.Error()))
		return err
	}
	err = client.SendCode(payload.TelephoneArea, payload.Telephone, payload.Content)

	if err != nil {
		logger.WithContext(ctx).Error("[SendSmsLogic] Send sms failed", logger.Field("error", err.Error()), logger.Field("payload", payload))
		createSms.Status = 2
	} else {
		createSms.Status = 1
	}
	logger.WithContext(ctx).Info("[SendSmsLogic] Send sms", logger.Field("telephone", payload.Telephone), logger.Field("content", createSms.Content))

	content, marshalErr = createSms.Marshal()
	if marshalErr != nil {
		return nil
	}
	audit.Content = string(content)
	if updateErr := l.deps.Store.Log().Update(ctx, audit); updateErr != nil {
		logger.WithContext(ctx).Error("[SendSmsLogic] Finalize sms log failed", logger.Field("error", updateErr.Error()), logger.Field("log_id", audit.Id))
	}
	return nil
}
