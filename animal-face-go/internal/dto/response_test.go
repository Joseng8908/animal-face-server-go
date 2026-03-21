package dto

import (
	"testing"
)

// Unit Test: Success Response 생성 확인
func TestNewSuccessResponse(t *testing.T) {
	data := "test_data"
	msg := "success"
	resp := NewSuccessResponse(data, msg)

	if !resp.Success {
		t.Errorf("Expected Success to be true, got %v", resp.Success)
	}

	if resp.Message != msg {
		t.Errorf("Expected Message to be %s, got %s", msg, resp.Message)
	}

	if resp.Data != data {
		t.Errorf("Expected Data to be %v, got %v", data, resp.Data)
	}
}
