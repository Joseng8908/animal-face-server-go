package middleware

import (
	"animal-face-go/internal/dto"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestErrorHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	// 미들웨어 등록
	r.Use(ErrorHandler())

	// 의도적으로 에러를 추가하는 핸들러
	r.GET("/error", func(c *gin.Context) {
		c.Error(assert.AnError) // 에러 강제 발생
		return
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/error", nil)
	r.ServeHTTP(w, req)

	// 검증
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response dto.APIResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.False(t, response.Success)
	assert.Equal(t, "Internal Server Error", response.Message)
}
