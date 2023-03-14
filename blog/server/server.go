package server

import (
	"chalet/blog/internal/controller/article"
	"chalet/blog/internal/controller/file"
	"chalet/blog/internal/controller/test"
	"chalet/blog/internal/controller/user"
	"chalet/blog/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.New()
	router.UseRawPath = true
	// 解决跨域
	router.Use(gin.Recovery(), middleware.CorsHandler())
	v1 := router.Group("/v1")
	initRouter(v1)
	router.Run(":8080")
}

// initRouter 初始化路由组
func initRouter(v1 *gin.RouterGroup) {
	//用户操作组
	rulesRoute := v1.Group("/user")
	initUserRouter(rulesRoute)
	//文章操作
	articleRoyte := v1.Group("/article")
	initArticleRouter(articleRoyte)
	//上传图片
	fileRoute := v1.Group("/file")
	FileRouter(fileRoute)

	testRoute := v1.Group("/test")
	TestRouter(testRoute)

}

func initUserRouter(rulesRoute *gin.RouterGroup) {
	{
		rulesRoute.POST("/register", user.Registration)
		rulesRoute.POST("/login", user.Login)
		//评论文章
	}
}

func initArticleRouter(rulesRoute *gin.RouterGroup) {
	{
		//上传文章
		rulesRoute.POST("/fondArticle", article.FondArticle)
		//倒叙获得文章列表
		rulesRoute.POST("/articleList", article.GetArticle)
		//点赞
		//收藏
	}
}

func FileRouter(rulesRoute *gin.RouterGroup) {
	{
		rulesRoute.POST("/image", file.UploadFile)
	}
}

func TestRouter(rulesRoute *gin.RouterGroup) {
	{
		rulesRoute.GET("/mail", test.SendMail)
	}
}
