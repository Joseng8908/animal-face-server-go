package dto

// 분석 결과 응답용 DTO
type AnimalUploadresponse struct {
	ID         uint    `json:"id`
	AnimalType string  `json:"animal_type"`
	Confidence float64 `json:"confidence"`
	ImageURL   string  `json:"image_url"`
}
