package domain

import "gorm.io/gorm"

// entity struct만들기
// Animalresult: 사용자의 얼굴 분석 결과를 담는 테이블

type AnimalResult struct {
	gorm.Model         // ID, CreatedAt, UPdateAT, DeleteAT을 자동으로 포함하는 엔티티 모델
	UserID     string  `json:"user_id", gorm:"not null`
	AnimalType string  `json:"animal_type" gorm:"not null` // 강아지상, 고양이상, 여웅상?
	Confidence float64 `json:"fonfidenc`                   // 정확도(0.0 ~ 1.0?) // 이건 들어봐야 알듯
	ImageURL   string  `json:"image_url`                   // 저장된 이미지 경로
}
