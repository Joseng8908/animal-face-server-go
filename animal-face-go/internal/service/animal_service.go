package service

import (
	"animal-face-go/internal/domain"
	"animal-face-go/internal/repository"
)

type AnimalService struct {
	repo *repository.AnimalRepository
}

func NewAnimalService(repo *repository.AnimalRepository) *AnimalService {
	return &AnimalService{repo: repo}
}

// UploadAndAnalyze
func (s *AnimalService) UploadAndAnalyze(userID string, filePath string) (*domain.AnimalResult, error) {
	// TODO: AI서버 호출
	// 지금은 mock데이터 넣기
	result := &domain.AnimalResult{
		UserID:     userID,
		AnimalType: "강아지상",
		Confidence: 0.88,
		ImageURL:   filePath,
	}

	err := s.repo.Save(result)
	return result, err
}
