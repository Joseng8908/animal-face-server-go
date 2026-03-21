package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCORSMiddleware(t *testing.T) {
	// 1. 테스트 설정
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(CORSMiddleware())

	r.GET("/test-cors", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// 2. 가짜 요청 생성 (OPTIONS 요청 테스트)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", "/test-cors", nil)

	// 3. 실행
	r.ServeHTTP(w, req)

	// 4. 검증: CORS 헤더가 포함되어 있는지 확인
	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "POST, OPTIONS, GET, PUT, DELETE", w.Header().Get("Access-Control-Allow-Methods"))
}
