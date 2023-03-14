package test

import (
	"chalet/blog/internal/service/test"
	"chalet/pkg/handlerutils"
	"github.com/gin-gonic/gin"
)

func SendMail(c *gin.Context) {
	if err := test.SendEmail(c); err != nil {
		handlerutils.FailWriteError(c, err)
		return
	}
	handlerutils.Ok(c)
}
