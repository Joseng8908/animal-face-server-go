package controller

import (
	"animal-face-go/internal/dto"
	"animal-face-go/internal/service"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AnimalController struct {
	service *service.AnimalService
}

func NewAnimalController(s *service.AnimalService) *AnimalController {
	return &AnimalController{service: s}
}

func (ctrl *AnimalController) UploadImage(c *gin.Context) {
	// 파일 받기
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("파일 업로드 실패", err.Error()))
		return
	}

	// 파일 저장 경로 설정(UUID로 이름 중복 방지하기)
	filename := uuid.New().String() + filepath.Ext(file.Filename)
	savePath := "uploads/" + filename

	// 실제 서버에 저장, gin제공 함수 사용
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("파일 서버에 저장 실패", err.Error()))
	}

	// 서비스 호출하기 DB저장 및 분석 요청
	userID := c.PostForm("user_id") // user_id를 폼 데이터로 받기
	result, err := ctrl.service.UploadAndAnalyze(userID, savePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("분석 결과 저장 실패", err.Error()))
	}

	// 성공 응답 보내기
	response := dto.NewSuccessResponse(dto.AnimalUploadresponse{
		ID:         result.ID,
		AnimalType: result.AnimalType,
		Confidence: result.Confidence,
		ImageURL:   result.ImageURL,
	}, "성공적으로 분석되었습니다.")
	c.JSON(http.StatusOK, response)
}
