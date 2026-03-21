package repository

import (
	"animal-face-go/internal/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	// 로컬 테스토용 sqlite사용, 나중에 다른 DB로 바꾸기
	DB, err = gorm.Open(sqlite.Open("animal_face.db"), &gorm.Config{})
	if err != nil {
		panic("데이터베이스 연결 실패")
	}

	DB.AutoMigrate(&domain.AnimalResult{})
}
