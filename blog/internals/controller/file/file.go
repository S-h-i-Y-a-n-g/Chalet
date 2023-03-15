package file

import (
	"chalet/pkg/handlerutils"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

//上传图片
func UploadFile(c *gin.Context) {
	// 先存储头像图片，获取Formdata的上传文件对象
	file, err := c.FormFile("image")
	// 判断是否有错误返回
	if err != nil {
		// 有则记录错误返回
		log.Printf("接收文件出错,%v", err)
		// 并直接返回结果
		handlerutils.FailWriteError(c, errors.New("上传图片失败"))
	}

	// 没有则获取到文件名,携带时间错拼成文件名
	fileName := strconv.Itoa(int(time.Now().Unix())) + "-" + file.Filename
	// 将目录和文件名组织在一起，形成保存路径
	savePath := "/Chalet/image" + "/" + fileName
	// 使用gin自带的保存方法，将文件保存到保存目录里。
	c.SaveUploadedFile(file, savePath)

	handlerutils.OkWithData(map[string]string{"image": "http://101.42.118.66:4030/" + fileName}, c)
}
