package redis

import (
	"chalet/dao/user"
	"chalet/pkg/entity"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis"
)


const (
	_UserUUID = "user_uuid"
)

var client = &redis.Client{}

func Init(redisConfig entity.RedisConfig) {
	client = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal("init redis failed:", err)
	}
}



//根据用户uid增加缓存
func SaveLoginUserInfo(ctx context.Context, UserUuId string, userInfo user.User) {
	//序列化用户信息
	UserInfoJSON, _ := json.Marshal(userInfo)
	// 前缀拼接key 存储userInfo 时长为1h,错误先抛出
	if err := client.Set(_UserUUID+UserUuId, string(UserInfoJSON), time.Hour).Err(); err != nil {
		log.Println(ctx, "redis.Set failed:", err)
	}
}

//查询缓存
func GetLoginUserInfo(ctx context.Context, UserUuId string) user.User {
	var userInfo user.User
	res, err := client.Get(_UserUUID + UserUuId).Result()
	if err != nil {
		log.Println("缓存查询失败")
		return userInfo
	}
	json.Unmarshal([]byte(res), &userInfo)

	return userInfo
}


//
