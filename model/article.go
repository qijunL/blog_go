package model

import (
	"time"
	)

//定义文章的结构体
type ArticleInfo struct {
	Id int64 `db:"id"`
	CategoryId  int64 `db:"id"`
	//文章的摘要
	Summary  string `db:"summary"`
	Title  string `db:"title"`
	ViewCount uint32 `db:"view_count"`
	CreateTime time.Time `db:"create_time"`
	CommentCount uint32 `db:"comment_count"`
	Username string `db:"username"`
}

type ArticleDetail struct {
	ArticleInfo

	//	文章内容
	Content string `db:"content"`
	Category

}

//用于文章上下页
type ArticleRecord struct {
	ArticleInfo
	Category
}