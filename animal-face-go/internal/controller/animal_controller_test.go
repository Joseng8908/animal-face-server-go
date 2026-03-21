package controller

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"animal-face-go/internal/domain"
	"animal-face-go/internal/repository"
	"animal-face-go/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUploadImage(t *testing.T) {
	// 1. 테스트 환경 설정 (DB -> Repo -> Service -> Controller)
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&domain.AnimalResult{})
	repo := repository.NewAnimalRepository(db)
	svc := service.NewAnimalService(repo)
	ctrl := NewAnimalController(svc)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/uploads", ctrl.UploadImage)

	// 2. 가짜 Multipart Form 데이터 생성 (핵심!)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// (1) 폼 필드 추가 (user_id)
	_ = writer.WriteField("user_id", "test_user_777")

	// (2) 파일 파트 추가 (image)
	part, _ := writer.CreateFormFile("image", "test_face.jpg")
	_, _ = io.WriteString(part, "fake-image-binary-data") // 가짜 데이터
	writer.Close()

	// 3. 요청 실행
	req, _ := http.NewRequest("POST", "/uploads", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// 4. 검증
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "성공적으로 분석되었습니다.")

	// 실제 파일이 uploads 폴더에 생성되었는지 확인 (선택사항)
	// 로직상 uuid로 생성되므로 폴더 내 파일 존재 여부만 체크 가능
}
