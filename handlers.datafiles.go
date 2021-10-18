package main

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"

	mycharts "github.com/stuartheriot/srcvisualiser/charts"
)

func showDataFiles(ctx *gin.Context) {
	if files, err := ioutil.ReadDir("./data/"); err == nil {
		render(
			ctx,
			gin.H{"title": "DataFiles to Chart", "payload": files},
			"datafiles.html",
		)
	} else { // silently fail
		render(
			ctx,
			gin.H{"title": "NO FILES FOUND", "payload": nil},
			"datafiles.html",
		)
	}
}

func chartFileData(ctx *gin.Context) {
	filename := ctx.Query("datafile")
	mycharts.ChartFileData(ctx, filename)
}
