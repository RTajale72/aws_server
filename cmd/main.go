package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.Get("/", Helloworld)
	r.Run()
}

func Helloworld(c *gin.Context) {
	c.json("Hello world")
}
