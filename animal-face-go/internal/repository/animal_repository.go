package repository

import (
	"animal-face-go/internal/domain"

	"gorm.io/gorm"
)

type AnimalRepository struct {
	db *gorm.DB
}

func NewAnimalRepository(db *gorm.DB) *AnimalRepository {
	return &AnimalRepository{db: db}
}

func (r *AnimalRepository) Save(result *domain.AnimalResult) error {
	return r.db.Create(result).Error
}

func (r *AnimalRepository) FindByUserID(userID string) ([]domain.AnimalResult, error) {
	var results []domain.AnimalResult
	err := r.db.Where("user_id = ?", userID).Find(&results).Error
	return results, err
}
