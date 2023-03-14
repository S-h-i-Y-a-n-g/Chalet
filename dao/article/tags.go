package article

import (
	"chalet/global"
	"chalet/infra/mysql"
	"context"
	"log"
)

type Tags struct {
	global.MODEL
	UserUUId string `json:"user_uu_id"`
	Name     string `json:"name"`
}

type TagArticle struct {
	TagId     int `json:"tag_id"`
	ArticleId int `json:"article_id"`
}

func (m *Tags) TableName() string {
	return "blog_tag"
}

func AddTagArticle(ctx context.Context, tagId int, articleId []int) {
	for _, v := range articleId {
		err := mysql.DB(ctx).Table("article_tag").Create(&TagArticle{TagId: tagId, ArticleId: v}).Error
		if err != nil {
			log.Fatalf("添加标签失败,文章id:%v", articleId)
		}
	}
}
