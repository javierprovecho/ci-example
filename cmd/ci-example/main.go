package main

import (
	"log"
	"os"

	"github.com/javierprovecho/ci-example/Godeps/_workspace/src/github.com/gin-gonic/gin"
)

const cats = false
const dogs = true

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	r := createRouter()

	r.Run(":" + port)
}

func createRouter() (r *gin.Engine) {
	r = gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hola Cylicon Valley! Estoy desde "+os.Getenv("enviroment"))
	})

	return
}
