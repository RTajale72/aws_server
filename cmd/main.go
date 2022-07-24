package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", Helloworld)
	r.Run()
}

func Helloworld(c *gin.Context) {
	c.JSON(http.StatusAccepted, "Hello world")
}
