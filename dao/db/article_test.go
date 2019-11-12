package db

import (
	"go_blog/model"
	"testing"
	"time"
)

func init() {
	//parseTime=true将mysql中的时间类型自动解析为go结构体中的时间类型
	dns := "root:admin@tcp(localhost:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

func TestInsertArticle(t *testing.T) {
	article := &model.ArticleDetail{}
	article.ArticleInfo.CategoryId = 1
	article.ArticleInfo.CommentCount = 0
	article.ArticleInfo.CreateTime = time.Now()
	article.ArticleInfo.Title = "大刷币"
	article.ArticleInfo.Username = "小明"
	article.ArticleInfo.Summary = "abc efd"
	article.ArticleInfo.ViewCount = 1
	article.Content = "wo ai wo jia"
	articleId, err := InsertArticle(article)
	if err != nil {
		return
	}
	t.Logf("articleId : %d\n", articleId)
}

func TestGetArticleList(t *testing.T) {
	articleList, err := GetArticleList(1, 15)
	if err != nil {
		return
	}
	t.Logf("articleList : %d\n", len(articleList))
}

func TestGetArticleListByCategoryId(t *testing.T) {

}
