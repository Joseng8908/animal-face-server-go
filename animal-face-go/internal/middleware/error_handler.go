package middleware

import (
	"animal-face-go/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 서버가 패닉이 나더라도 서버가 죽지 않고, 에러메세지를 던지도록 보강
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // 다음 핸들러 실행

		// 핸들러 실행 중 에러 발생 확인
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("Internal Server Error", err.Error()))
		}
	}
}
