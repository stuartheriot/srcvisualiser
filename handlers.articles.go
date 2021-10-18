package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(ctx *gin.Context) {
	articles := getAllArticles()

	render(
		ctx,
		gin.H{"title": "SrcViz", "payload": articles},
		"index.html",
	)
}

func getArticle(ctx *gin.Context) {
	if articleID, err := strconv.Atoi(ctx.Param("article_id")); err == nil {
		if article, err := getArticleByID(articleID); err == nil {
			render(
				ctx,
				gin.H{"title": article.Title, "payload": article},
				"article.html",
			)
		}
	}
}
