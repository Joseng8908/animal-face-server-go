package main

import (
	"animal-face-go/internal/dto"
	"animal-face-go/internal/middleware"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New() // New를 사용해서 gin을 내 입맛대로 커스터마이징
	// 커스텀마이징, 미들웨어 등록하기
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.ErrorHandler())

	r.GET("/health", func(c *gin.Context) {
		response := dto.NewSuccessResponse(nil, "서버가 정상입니다")
		c.JSON(http.StatusOK, response)
	})

	r.Run(":8080")
}
