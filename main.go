package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//测试
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "%s", "ok")
	})
	log.Fatal(r.Run(":8080"))
}
