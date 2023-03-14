package test

import (
	"chalet/blog/internal/config"
	"chalet/pkg/mail"
	"context"
)

func SendEmail(c context.Context) error {
	//, "3062455224@qq.com"
	return mail.SendMail(config.AppConfig.QQMail, []string{"844310197@qq.com"}, "紧急通知！！！", "云顶之弈都不玩了？")
}
