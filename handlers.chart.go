package main

import (
	"github.com/gin-gonic/gin"

	"github.com/stuartheriot/srcvisualiser/charts"
)

func showChart(ctx *gin.Context) {
	charts.BuildChart(ctx)
}
