package main

import (
	"animal-face-go/internal/middleware"

	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	r.GET("/health")

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
}
