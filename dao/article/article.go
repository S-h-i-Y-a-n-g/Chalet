package article

import (
	"chalet/global"
	"chalet/infra/mysql"
	"context"
)

type Article struct {
	global.MODEL
	UserUuid   string `json:"user_uuid"`
	Title      string `json:"title"`
	Context    string `json:"context"`
	Category   int    `json:"category"`
	Views      int
	Collection int
}

func (m *Article) TableName() string {
	return "blog_article"
}

func FondArticle(ctx context.Context, article Article) (error, uint) {
	return mysql.DB(ctx).Save(&article).Error, article.ID
}

func GetArticle(ctx context.Context, start, end int) ([]Article, error) {
	var articleArr []Article
	err := mysql.DB(ctx).Debug().Order("id desc").Offset(start).Limit(end).Find(&articleArr).Error
	if err != nil {
		return nil, err
	}
	return articleArr, nil
}
