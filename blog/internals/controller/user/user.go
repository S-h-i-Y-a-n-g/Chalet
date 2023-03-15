package user

import (
	"chalet/blog/internals/service/user"
	"chalet/pkg/handlerutils"
	"chalet/pkg/request"
	"github.com/gin-gonic/gin"
)

// Registration user
func Registration(c *gin.Context) {
	var newUser request.RegistrationReq
	if err := handlerutils.BindAndValidate(c, &newUser); err != nil {
		handlerutils.WriteError(c, err)
		return
	}
	if err := user.Registration(c, newUser); err != nil {
		handlerutils.WriteError(c, err)
		return
	}
	handlerutils.Ok(c)
}

// Login 用户登录
func Login(c *gin.Context) {
	var LoginUser request.LoginReq
	if err := handlerutils.BindAndValidate(c, &LoginUser); err != nil {
		return
	}
	userInfo, err := user.Login(c, LoginUser)
	if err != nil {
		handlerutils.WriteError(c, err)
		return
	}
	handlerutils.OkWithData(userInfo, c)
}

//删除用户

//修改用户密码

//获取用户信息

//设置用户信息
