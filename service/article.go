package service

import (
	"go_blog/dao/db"
	"go_blog/model"
)

//获取文章和对应的分类

func GetArticleRecordList(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	//	1获取文章的列表
	articleInfoList, err := db.GetArticleList(pageNum, pageSize)
	if err != nil {
		return
	}
	if len(articleInfoList) <= 0 {
		return
	}

	//获取文章对应的分类（多个）
	categoryIds := getCategoryIds(articleInfoList)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		return
	}
	//返回页面做聚合
	//遍历所有的文章
	for _, article := range articleInfoList {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		//文章取出分类id
		categoryId := article.CategoryId
		//遍历分类列表
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

//根据文章的id获取多个分类的ID
func getCategoryIds(articleInfoList []*model.ArticleInfo) (ids []int64) {
	//	遍历文章得到每一个文章
LABEL:
	for _, article := range articleInfoList {
		//	从当前文章取出分类ID
		categoryId := article.CategoryId
		//去重防止重复
		for _, id := range ids {
			if id == categoryId {
				continue LABEL
			}
		}
		ids = append(ids, categoryId)
	}
	return
}

//根据分类id，获取该类文章和他们对应的分类信息
func GetArticleRecordListById(categoryId, pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	//	1获取文章的列表
	articleInfoList, err := db.GetArticleListByCategoryId(categoryId, pageNum, pageSize)
	if err != nil {
		return
	}
	if len(articleInfoList) <= 0 {
		return
	}

	//获取文章对应的分类（多个）
	categoryIds := getCategoryIds(articleInfoList)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		return
	}
	//返回页面做聚合
	//遍历所有的文章
	for _, article := range articleInfoList {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		//文章取出分类id
		categoryId := article.CategoryId
		//遍历分类列表
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}
