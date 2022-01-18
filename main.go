package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHostPort() (string, int) {
	return "127.0.0.1", 8000
}

func sayAliveView(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{})
}

func getRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/:key", sayAliveView)
	return router
}

func main() {
	host, port := getHostPort()
	router := getRouter()
	router.Run(fmt.Sprintf("%s:%d", host, port))
}
