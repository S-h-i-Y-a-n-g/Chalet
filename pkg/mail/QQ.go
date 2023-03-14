package mail

import (
	"chalet/pkg/entity"
	"crypto/tls"
	"errors"
	"fmt"
	"gopkg.in/gomail.v2"
)

func SendMail(mail entity.QQMail, mailTo []string, subject, body string) error {
	//创建一个Message实例，Message提供了整个邮件协议内容的构建，默认实例采用UTF-8字符集和Quoted-printable编码
	//使用其他编码格式 gomail.SetEncoding(gomail.Base64)
	m := gomail.NewMessage()

	//发送人
	// m.SetHeader("From", "test"+"<"+userName+">") // 增加发件人别名
	m.SetHeader("From", mail.Nickname+"<"+mail.UserName+">")

	//接收人
	m.SetHeader("To", mailTo...)

	//抄送人
	//m.SetAddressHeader("Cc", "xxx@qq.com", "xiaozhujiao")
	//m.SetHeader("Cc", "******@qq.com")                  // 抄送，可以多个
	//m.SetHeader("Bcc", "******@qq.com")                 // 暗送，可以多个

	//主题
	m.SetHeader("Subject", subject)
	//内容
	// text/html 的意思是将文件的 content-type 设置为 text/html 的形式，浏览器在获取到这种文件时会自动调用html的解析器对文件进行相应的处理。
	// 可以通过 text/html 处理文本格式进行特殊处理，如换行、缩进、加粗等等
	m.SetBody("text/html", body)

	// text/plain的意思是将文件设置为纯文本的形式，浏览器在获取到这种文件时并不会对其进行处理
	// m.SetBody("text/plain", "纯文本")

	//附件文件，可以是文件，照片，视频等等
	//m.Attach("./myIpPic.png")

	//拿到token，并进行连接,第4个参数是填授权码 AuthorizationCode=qpqtchcyqnvhbegd
	d := gomail.NewDialer(mail.Host, mail.Port, mail.UserName, mail.AuthorizationCode)
	//可以通过gomail.Rename设置附件别名，mime.QEncoding.Encode()设置格式为UTF-8，防止乱码。
	//m.Attach(url, gomail.Rename(path.Base(file_name)))==>m.Attach(url, gomail.Rename(mime.QEncoding.Encode(path.Base(file_name))))

	// 设置为true时，关闭TLS认证，否则默认开启，需要配置证书认证，稍微麻烦一些
	d.TLSConfig = &tls.Config{InsecureSkipVerify: mail.IsSsl}

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		return errors.New(fmt.Sprintf("DialAndSend err %v:", err))
	}
	fmt.Printf("send mail success\n")
	return nil
}
