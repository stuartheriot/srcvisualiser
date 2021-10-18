package main

import "github.com/gin-gonic/gin"

func initializeRoutes(router gin.IRouter) {
	router.GET("/srcviz", showIndexPage)

	router.GET("/article/view/:article_id", getArticle)

	router.GET("/chart", showDataFiles)

	router.GET("/chart/display", chartFileData)
}
