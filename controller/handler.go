package controller

import (
	"github.com/gin-gonic/gin"
	"go_blog/service"
	"net/http"
	"strconv"
)

//访问主页的控制器

func IndexHandle(c *gin.Context) {
	//从service中取
	articleRecordList, err := service.GetArticleRecordList(0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	//加载分类数据
	categoryList, err := service.GetALLCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	//第一种方式
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  articleRecordList,
		"category_list": categoryList,
	})
	//gin.H本质上是一个map
	//第二种方式
	//var data map[string]interface{} = make(map[string]interface{},16)
	//data["article_list"]= articleRecordList
	//data["category_list"] = categoryList
	//c.HTML(http.StatusOK,"views/index.html",data)

}

//点击分类云进行分类
func CategoryList(c *gin.Context) {
	CategoryIdStr := c.Query("category_id")
	//转
	categoryId, err := strconv.ParseInt(CategoryIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	//根据分类id获取文章列表
	articleRecordList, err := service.GetArticleRecordListById(int(categoryId), 0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	//再次加载分类数据，用于分类云显示
	categoryList, err := service.GetALLCategoryList()
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  articleRecordList,
		"category_list": categoryList,
	})
}
