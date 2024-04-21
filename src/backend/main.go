package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type linkJson struct {
	Start      string `json:"start"`
	End        string `json:"end"`
	IsMultiple bool   `json:"isMultiple"`
}

var baseURL = "https://en.wikipedia.org"

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	r.POST("/multiple/bfs", func(c *gin.Context) {
		var link linkJson

		if err := c.BindJSON(&link); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		start := link.Start
		end := link.End

		result := bfs(start, end, baseURL)

		c.JSON(http.StatusOK, result)
	})

	r.POST("/single/bfs", func(c *gin.Context) {
		var link linkJson

		if err := c.BindJSON(&link); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		start := link.Start
		end := link.End

		result := bfs_single(start, end, baseURL)

		c.JSON(http.StatusOK, result)
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
