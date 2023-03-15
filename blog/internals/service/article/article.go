package article

import (
	"chalet/dao/article"
	"chalet/pkg/request"
	"context"
	"errors"
)

func FondArticleHandle(ctx context.Context, newArticle request.FondArticleReq) (error, int) {
	//通过生成的文章id，生成redis的浏览人数和收藏量
	err, articleId := article.FondArticle(ctx, BuildArticle(newArticle))
	if err != nil {
		return errors.New("FondArticleHandle fail"), 0
	}
	article.AddTagArticle(ctx, int(articleId), newArticle.Tag)

	return nil, int(articleId)
}

// BuildArticle 构建插入文章&tag
func BuildArticle(newArticle request.FondArticleReq) (article article.Article) {
	article.UserUuid = newArticle.UserUuid
	article.Title = newArticle.Title
	article.Context = newArticle.Context
	article.Category = newArticle.Category
	return
}

func GetArticleHandle(ctx context.Context, start, end int) ([]article.Article, error) {
	//通过生成的文章id，生成redis的浏览人数和收藏量
	if start == end && end == 0 {
		start = 0
		end = 10
	}
	Array, err := article.GetArticle(ctx, start, end)
	if err != nil {
		return nil, errors.New("GetArticleHandle fail")
	}
	return Array, nil
}
