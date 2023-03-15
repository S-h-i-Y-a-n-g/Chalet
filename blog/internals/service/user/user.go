package user

import (
	"chalet/dao/user"
	"chalet/infra/redis"
	"chalet/pkg/request"
	"context"
)

func Registration(ctx context.Context, newUser request.RegistrationReq) error {
	//md5简单加密一下
	//TODO 注册成功返回token直接用于登录
	return user.Registration(ctx, newUser)
}

func Login(ctx context.Context, User request.LoginReq) (user.User, error) {
	// TODO 先从缓存查
	redis.GetLoginUserInfo(ctx,User.Username)
	//缓存没有则进mysql，并加个半小时的缓存
	UserInfo, err := user.Login(ctx, User)

	return UserInfo, err
}



