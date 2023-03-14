package user

import (
	"chalet/global"
	"chalet/infra/mysql"
	"chalet/pkg/request"
	"chalet/pkg/utils"
	"context"
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	global.MODEL
	Signature       string    `json:"signature"`
	UUID            uuid.UUID `json:"uuid" gorm:"not null;comment:用户UUID"`
	Username        string    `json:"username"`
	Password        string    `json:"password"`
	NickName        string    `json:"nick_name"`
	SideMode        string    `json:"side_mode"`
	HeaderImg       string    `json:"header_img"`
	BaseColor       string    `json:"base_color"`
	ActiveColor     string    `json:"active_color"`
	AuthorityId     int       `json:"authority_id"`
	Phone           string    `json:"phone"`
	Email           string    `json:"email"`
	Enable          int       `json:"enable"`
	BackgroundImage string    `json:"background_image"`
}

func TableName() string {
	return "blog_user"
}

func Registration(ctx context.Context, newUser request.RegistrationReq) error {
	var user User
	//组装一下默认数值
	if !errors.Is(mysql.DB(ctx).Table(TableName()).Where("username = ?", newUser.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("邮箱已注册，请登录")
	}

	user.UUID = uuid.NewV4()
	user.Password = utils.MD5V([]byte(newUser.Password))
	user.Username = newUser.Username
	user.Email = newUser.Email
	user.NickName = newUser.Nickname
	user.HeaderImg = newUser.HeaderImg
	user.Enable = 1
	return mysql.DB(ctx).Table(TableName()).Create(&user).Error
}

func Login(ctx context.Context, newUser request.LoginReq) (User, error) {
	var user User
	//组装一下默认数值
	if errors.Is(mysql.DB(ctx).Table(TableName()).Where("username = ?", newUser.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是存在
		return user, errors.New("user not exit")
	}
	//判断密码
	if user.Password != utils.MD5V([]byte(newUser.Password)) {
		return user, errors.New("password error")
	}
	return user, nil
}
