package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//backend_port := os.Getenv("BACKEND_PORT")
	r := gin.Default()
	r.GET("/hello", Helloworld)
	r.Run(":8000")
}

func Helloworld(c *gin.Context) {
	c.JSON(http.StatusAccepted, "Hello world")
}
