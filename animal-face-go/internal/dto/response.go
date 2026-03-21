package dto

// APIResponse 는 모든 API의 공통 응답 규격
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SuccessResponse: 성공시 응답 객체
func NewSuccessResponse(data interface{}, message string) APIResponse {
	return APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

// ErrorResponse: 실패시 응답 객체
func NewErrorResponse(message string, err string) APIResponse {
	return APIResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
