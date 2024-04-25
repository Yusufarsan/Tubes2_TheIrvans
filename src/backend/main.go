package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type linkJson struct {
	Start string `json:"start"`
	End   string `json:"end"`
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

		time_start := time.Now()
		result, visitedURLCount := bfs(start, end, baseURL)
		time_elapsed := time.Since(time_start)

		c.JSON(http.StatusOK, gin.H{"result": result, "articles_count": visitedURLCount, "time_elapsed": time_elapsed.Milliseconds()})
	})

	r.POST("/single/bfs", func(c *gin.Context) {
		var link linkJson

		if err := c.BindJSON(&link); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		start := link.Start
		end := link.End

		time_start := time.Now()
		result, visitedURLCount := bfs_single(start, end, baseURL)
		time_elapsed := time.Since(time_start)

		c.JSON(http.StatusOK, gin.H{"result": result, "articles_count": visitedURLCount, "time_elapsed": time_elapsed.Milliseconds()})
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
