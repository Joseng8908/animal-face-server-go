package repository

import (
	"animal-face-go/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestAnimalRepository(t *testing.T) {
	// 1. 테스트용 인메모리 DB 설정
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&domain.AnimalResult{})
	repo := NewAnimalRepository(db)

	// 2. 데이터 저장 테스트
	testResult := &domain.AnimalResult{
		UserID:     "user123",
		AnimalType: "강아지상",
		Confidence: 0.95,
	}

	err := repo.Save(testResult)
	assert.NoError(t, err)

	// 3. 데이터 조회 테스트
	found, err := repo.FindByUserID("user123")
	assert.NoError(t, err)
	assert.Len(t, found, 1)
	assert.Equal(t, "강아지상", found[0].AnimalType)
}
