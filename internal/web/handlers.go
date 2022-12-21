package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func homeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func priceChartHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"x": priceChart.X,
		"y": priceChart.Y,
	})
}
