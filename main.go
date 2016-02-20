package main

import "github.com/gin-gonic/gin"

const cats = false
const dogs = true

func main() {
	r := createRouter()

	r.Run(":8080")
}

func createRouter() (r *gin.Engine) {
	r = gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello world!")
	})

	return
}
