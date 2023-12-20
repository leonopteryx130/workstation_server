package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print(888)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		fmt.Println("received")
		c.String(http.StatusOK, "hello World!")
	})

	r.Run(":7979")

	//test
}
