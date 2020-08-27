package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/html/*")
	r.Static("/static", "./static")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title": "EduHub",
		})
	})

	r.Run(":9000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
