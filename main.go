package main

import (
	"github.com/gin-gonic/gin"
	"go_blog/controller"
	"go_blog/dao/db"
)

func main() {
	router := gin.Default()
	dns := "root:admin@tcp(localhost:3306)/blogger?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}
	//加载静态文件
	router.Static("/static/", "./static/")
	//加载模板
	router.LoadHTMLGlob("views/*")
	router.GET("/", controller.IndexHandle)
	router.GET("/category/")
	_ = router.Run(":8000")
}
