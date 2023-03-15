package article

import (
	"chalet/blog/internals/service/article"
	"chalet/pkg/handlerutils"
	"chalet/pkg/request"
	"github.com/gin-gonic/gin"
)

// FondArticle 上传文章
func FondArticle(c *gin.Context) {
	var newArticle request.FondArticleReq
	if err := handlerutils.BindAndValidate(c, &newArticle); err != nil {
		handlerutils.WriteError(c, err)
		return
	}
	err, data := article.FondArticleHandle(c, newArticle)
	if err != nil {
		handlerutils.WriteError(c, err)
		return
	}
	handlerutils.OkWithData(data, c)
}

//倒叙获得文章展示
func GetArticle(c *gin.Context) {
	var req request.GetArticleReq
	if err := handlerutils.BindAndValidate(c, &req); err != nil {
		handlerutils.WriteError(c, err)
		return
	}
	// 查一下登录状态
	data, err := article.GetArticleHandle(c, req.Page*req.Size, req.Page*req.Size+req.Size)
	if err != nil {
		handlerutils.WriteError(c, err)
		return
	}
	handlerutils.OkWithData(data, c)
}
