package main

import (
	"fmt"

	"github.com/aoimaru/gws/lib"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(lib.Test())
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
