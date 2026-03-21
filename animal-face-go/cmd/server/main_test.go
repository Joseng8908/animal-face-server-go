package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"animal-face-go/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// 1. 테스트용 Gin 엔진 설정
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, dto.NewSuccessResponse(nil, "서버가 정상입니다."))
	})

	// 2. 가짜 요청(Mock Request) 생성
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)

	// 3. 요청 실행
	r.ServeHTTP(w, req)

	// 4. 검증 (Assert)
	assert.Equal(t, http.StatusOK, w.Code)

	var response dto.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.NoError(t, err)
	assert.True(t, response.Success)
	assert.Equal(t, "서버가 정상입니다.", response.Message)
}
