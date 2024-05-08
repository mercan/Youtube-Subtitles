package utils

type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewErrorResponse(message string) BaseResponse {
	return BaseResponse{
		Success: false,
		Message: message,
	}
}

func NewSuccessResponse(data interface{}) BaseResponse {
	return BaseResponse{
		Success: true,
		Data:    data,
	}
}
